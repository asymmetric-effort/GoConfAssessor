// file: cmd/verify-manifest/main.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>
package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/manifest"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	var (
		err          error
		output       []byte
		rootManifest manifest.Manifest
	)
	// Define flags
	debug := flag.Bool("debug", false, "Enable debug-level logging")
	manifestFile := flag.String("manifest", "", "Path to root YAML manifest (required)")
	flag.Parse()

	// Verify required flag
	if *manifestFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	log := logger.Logger
	if err = logger.SetLevel(logger.IsDebug(*debug)); err != nil {
		log.Error(err)
	}
	log.Debug("logger ready")

	// Load and resolve the manifest tree
	log.Debug("loading manifest")
	if err = rootManifest.Load(*manifestFile); err != nil {
		log.Fatalf("Manifest(%q) failure: %v", *manifestFile, err)
	}

	log.Debug("marshall manifest")
	if output, err = yaml.Marshal(rootManifest); err != nil {
		log.Fatal(err)
	}
	// For now, assume manifest.Load performs all syntax checks.
	fmt.Printf(string(output))
	log.Debug("Manifest successfully loaded and resolved")
}
