// file: pkg/manifest/Assertion.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// Assertion represents a single check in a group.
type Assertion struct {
	Label string `yaml:"label,omitempty"`

	Statement string `yaml:"statement,omitempty"`

	Expected struct {
		Type string `yaml:"type,omitempty"`

		Value interface{} `yaml:"value,omitempty"`
	} `yaml:"expected"`

	Weight uint `yaml:"weight,omitempty"`

	Source struct {
		Path string `yaml:"path,omitempty"`

		Pattern string `yaml:"pattern,omitempty"`
	} `yaml:"source,omitempty"`

	Operator string `yaml:"operator,omitempty"`
}
