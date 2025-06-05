// file: pkg/manifest/RootManifest.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// RootManifest represents the fully resolved manifest with general metadata and assertion groups.
type RootManifest struct {
	// The General section is required only in the root manifest
	// This is mostly metadata, versioning and other information used to organize information.
	General GeneralSection `yaml:"general:omitempty"`
	// Facts are data which can be defined in files and referenced by assertions.
	// For example, a "fact" might be a list of dns servers assertions can reference
	// to validate that the analyzed config file contains these facts.
	Facts []FactDescriptor
	// Patterns can be defined in any file.  These are reusable global-scope items
	Patterns []PatternDescriptor `yaml:"patterns,omitempty"`
	// Assertions are groups of factual statements evaluated for truthiness
	Assertions []AssertionGroup `yaml:"assertions,omitempty""`
}
