// file: pkg/manifest/PatternDescriptor.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// PatternDescriptor - A pattern descriptor describes a reusable regular expression
// this can be referenced later using the pattern::<name> format
type PatternDescriptor struct {
	Include string `yaml:"include,omitempty"`

	Pattern struct {
		Name string `yaml:"name"`

		Regex string `yaml:"regex,omitempty"`
	} `yaml:"pattern"`
}
