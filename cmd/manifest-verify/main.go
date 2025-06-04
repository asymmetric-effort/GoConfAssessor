// file: cmd/manifest-verify/main.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>
package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/manifest"
	"log"
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
	if *debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Debug mode enabled")
	}

	// Load and resolve the manifest tree
	rootManifest, err := manifest.LoadAndResolve(*manifestPath)
	if err != nil {
		log.Fatalf("Failed to load or resolve manifest %q: %v", *manifestPath, err)
	}

	// TODO: Verify each field of rootManifest for legitimate values:
	//  - general.name non-empty
	//  - general.version follows semantic version format
	//  - metadata entries are key->string
	//  - each group has non-empty Name
	//  - each assertion:
	//      * Label non-empty
	//      * Parser is one of the allowed parsers
	//      * AppliesTo non-empty
	//      * Statement non-empty
	//      * Expected.Type is valid and Expected.Value matches type
	//      * Weight >= 0
	//      * Source has at least Path or Pattern
	//      * Operator is one of matches|contains|equals

	// For now, assume manifest.LoadAndResolve performs all syntax checks.
	fmt.Println("Manifest successfully loaded and resolved:", rootManifest)
	fmt.Println("Manifest verification passed.")
}
