// file: cmd/manifest-verify/main.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>
package main

import (
	"flag"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/manifest"
	"os"
)

func main() {
	// Define flags
	debug := flag.Bool("debug", false, "Enable debug-level logging")
	manifestPath := flag.String("manifest", "", "Path to root YAML manifest (required)")
	flag.Parse()

	// Verify required flag
	if *manifestPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	log := logger.Logger
	if *debug {
		if err := logger.SetLevel("debug"); err != nil {
			log.Fatalf("error setting log level: %v", err)
		}
	} else {
		if err := logger.SetLevel("info"); err != nil {
			log.Fatalf("error setting log level: %v", err)
		}
	}
	log.Debug("logger ready")

	var (
		rootManifest manifest.Manifest
		err          error
	)

	// Load and resolve the manifest tree
	if err = rootManifest.Load(*manifestPath); err != nil {
		log.Fatalf("Manifest(%q) failure: %v", *manifestPath, err)
	}

	// For now, assume manifest.Load performs all syntax checks.
	log.Info("Manifest successfully loaded and resolved:", rootManifest)
}
