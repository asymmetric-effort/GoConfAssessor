package manifest

// FactDescriptor - describes facts, obviously
type FactDescriptor struct {
	Include string `yaml:"include,omitempty"`

	Fact struct {
		Name string `yaml:"name"`

		Data interface{} `yaml:"data,omitempty"`
	} `yaml:"fact"`
}
