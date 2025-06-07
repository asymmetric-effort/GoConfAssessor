// file: cmd/assess/main.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>
package main

import "fmt"

/**
 * This program performs an assessment of a directory of source config files
 * against a YAML manifest tree of assertions and produces a YAML report file
 * attesting to the outcome of the assessment.
 */
func main() {
	fmt.Print("Assessment starting")
	// ToDo: load the manifest tree
	// ToDo: validate the manifest tree
	// ToDo: for every config file in the source directory, compare to the manifest tree
}
