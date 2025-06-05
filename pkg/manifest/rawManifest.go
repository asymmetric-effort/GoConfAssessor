// file: pkg/manifest/rawManifest.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// rawManifest - raw structures for unmarshaling
type rawManifest struct {
	General    GeneralSection           `yaml:"general,omitempty"`
	Assertions []rawAssertionGroupEntry `yaml:"assertions"`
}
