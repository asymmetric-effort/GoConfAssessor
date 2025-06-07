// file: cmd/assess/main.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>
package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/manifest"
	"github.com/sam-caldwell/GoConfAssessor/pkg/report"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
	"os"
	"path/filepath"
)

/**
 * This program performs an assessment of a directory of source config files
 * against a YAML manifest tree of assertions and produces a YAML report file
 * attesting to the outcome of the assessment.
 */
func main() {
	fmt.Print("Assessment starting")
	var (
		err          error
		log          = logger.Logger
		rootManifest manifest.Manifest
	)
	// Define flags
	debug := flag.Bool("debug", false, "Enable debug-level logging")
	manifestFile := flag.String("manifest", "", "Path to root YAML manifest (required)")
	sourceDir := flag.String("source", "", "Source directory containing files to assess")
	reportDir := flag.String("report", "", "Directory where reports will be written")
	confType := flag.String("format", "", "Configuration file type (e.g. cisco ios)")
	flag.Parse()
	if err = utils.VerifyArgs(log, debug, manifestFile, sourceDir, reportDir, confType); err != nil {
		log.Fatal(err)
	}
	if err = rootManifest.Load(*manifestFile); err != nil {
		log.Error(err)
	}
	//parser := parsers.New(confType)
	err = filepath.Walk(*sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			// skip files or directories we can't stat or just directories in general.
			return nil
		}
		//data, err := os.ReadFile(path)
		//if err != nil {
		//	return nil
		//}

		report.Start(*reportDir)
		if err = report.Write(); err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success")
}
