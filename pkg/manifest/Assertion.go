// file: pkg/manifest/Assertion.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// Assertion represents a single check in a group.
type Assertion struct {
	Label string `yaml:"label,omitempty"`

	Statement string `yaml:"statement,omitempty"`

	Expected Expectation `yaml:"expected"`

	Weight uint `yaml:"weight,omitempty"`

	Source ActualSource `yaml:"source,omitempty"`

	Operator string `yaml:"operator,omitempty"`
}
