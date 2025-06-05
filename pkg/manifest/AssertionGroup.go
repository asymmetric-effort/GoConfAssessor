// file: pkg/manifest/AssertionGroup.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// AssertionGroup represents a named collection of assertions.
type AssertionGroup struct {
	Include string `yaml:"include,omitempty"`

	Name string `yaml:"name,omitempty"`

	Parser string `yaml:"parser,omitempty"`

	Items []Assertion `yaml:"items,omitempty"`
}
