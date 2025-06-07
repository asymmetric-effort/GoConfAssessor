// file: cmd/create-manifest/main.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>
package main

import (
	"flag"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/manifest"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func main() {
	// Define flags
	debug := flag.Bool("debug", false, "Enable debug-level logging")
	manifestPath := flag.String("manifest", "", "Path to root YAML manifest (required)")
	name := flag.String("name", "", "The name of the test manifest tree")
	author := flag.String("author", utils.GetCurrentUsername(), "The manifest-tree author")
	flag.Parse()

	// Verify required flags
	if *manifestPath == "" || strings.TrimSpace(*name) == "" {
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
	if err = rootManifest.Create(manifestPath, name, author); err != nil {
		log.Fatalf("Manifest(%q) failure: %v", *manifestPath, err)
	}

	var output []byte
	if output, err = yaml.Marshal(rootManifest); err != nil {
		log.Fatal(err)
	}
	// For now, assume manifest.Load performs all syntax checks.
	if err = os.WriteFile(*manifestPath, []byte(output), 0644); err != nil {
		log.Fatal(err)
	}

	log.Debug("Manifest successfully loaded and resolved")
}
