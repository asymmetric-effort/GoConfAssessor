// file: pkg/manifest/rawAssertionGroupEntry.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// rawAssertionGroupEntry - the raw yaml assertion group subsection
type rawAssertionGroupEntry struct {
	Include string             `yaml:"include"`
	Name    string             `yaml:"name"`
	Items   []rawAssertionItem `yaml:"items"`
}
