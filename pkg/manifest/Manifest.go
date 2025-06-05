// file: pkg/manifest/Manifest.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// Manifest represents the fully resolved manifest with general metadata and assertion groups.
type Manifest struct {
	//
	// The General section is required only in the root manifest
	// This is mostly metadata, versioning and other information used to organize information.
	General GeneralSection `yaml:"general:omitempty"`
	//
	// Facts are data which can be defined in files and referenced by assertions.
	// For example, a "fact" might be a list of dns servers assertions can reference
	// to validate that the analyzed config file contains these facts.
	Facts []FactDescriptor `yaml:"facts,omitempty"`
	//
	// Patterns can be defined in any file.  These are reusable global-scope items
	// A pattern is a regular expression used by assertions to match parts of an
	// analyzed configuration.
	Patterns []PatternDescriptor `yaml:"patterns,omitempty"`
	//
	// Assertions are groups of statements which can be evaluated against an analyzed configuration
	// using facts and patterns.  Assertions must be true, or they will fail, and reports from assessing
	// these assertions record the pass/fail state of the assertion.
	Assertions []AssertionGroup `yaml:"assertions,omitempty""`
}
