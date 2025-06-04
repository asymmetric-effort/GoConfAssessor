// file: cmd/assess/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sam-caldwell/GoConfAssessor/pkg/evaluator"
	"github.com/sam-caldwell/GoConfAssessor/pkg/manifest"
	"github.com/sam-caldwell/GoConfAssessor/pkg/report"
)

func main() {
	// Define flags
	debug := flag.Bool("debug", false, "Enable debug‐level logging")
	manifestPath := flag.String("manifest", "", "Path to root YAML manifest (required)")
	sourceDir := flag.String("source", "", "Directory of config files to assess (required)")
	reportDir := flag.String("report-dir", "", "Directory to write YAML reports (required)")
	flag.Parse()

	// Verify required flags
	if *manifestPath == "" || *sourceDir == "" || *reportDir == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Debug mode enabled")
	}

	// 1) Load and preprocess the manifest (resolving includes, prefixing labels, etc.)
	rootManifest, err := manifest.LoadAndResolve(*manifestPath)
	if err != nil {
		log.Fatalf("Failed to load manifest %q: %v", *manifestPath, err)
	}
	if *debug {
		log.Printf("Loaded manifest: %+v\n", rootManifest)
	}

	// 2) Enumerate all files under sourceDir
	configFiles := []string{}
	err = filepath.Walk(*sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			configFiles = append(configFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to scan source directory %q: %v", *sourceDir, err)
	}
	if len(configFiles) == 0 {
		log.Fatalf("No files found under source directory %q", *sourceDir)
	}
	if *debug {
		log.Printf("Found %d config files under %s\n", len(configFiles), *sourceDir)
	}

	// 3) For each config file:
	for _, cfgPath := range configFiles {
		if *debug {
			log.Printf("Processing %s\n", cfgPath)
		}
		// 3a) Choose parser based on file extension or manifest-level override
		parserType := manifest.DetermineParser(rootManifest, cfgPath)
		deviceCfg, err := evaluator.LoadDeviceConfig(cfgPath, parserType)
		if err != nil {
			log.Printf("Skipping %s: failed to load device config: %v\n", cfgPath, err)
			continue
		}

		// 3b) Evaluate all assertions against this deviceConfig
		results, err := evaluator.RunAll(rootManifest.Assertions, deviceCfg)
		if err != nil {
			log.Printf("Evaluation error for %s: %v\n", cfgPath, err)
			continue
		}

		// 3c) Write a YAML report named <basename>-report.yaml under reportDir
		base := filepath.Base(cfgPath)
		reportName := fmt.Sprintf("%s-report.yaml", base)
		outPath := filepath.Join(*reportDir, reportName)
		if err := report.Write(outPath, rootManifest.General, results); err != nil {
			log.Printf("Failed to write report for %s: %v\n", cfgPath, err)
			continue
		}
		if *debug {
			log.Printf("Wrote report: %s\n", outPath)
		}
	}
}
