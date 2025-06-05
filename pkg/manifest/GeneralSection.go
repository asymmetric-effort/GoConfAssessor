// file: pkg/manifest/GeneralSection.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// GeneralSection - The 'general' section of a YAML manifest file
type GeneralSection struct {
	Name     string            `yaml:"name"`
	Version  string            `yaml:"version"`
	Metadata map[string]string `yaml:"metadata,omitempty"`
}
