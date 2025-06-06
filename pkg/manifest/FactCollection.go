package manifest

// FactCollection - describes facts, obviously
type FactCollection struct {
	Include string `yaml:"include,omitempty"`

	Fact string      `yaml:"fact"`
	Data interface{} `yaml:"data,omitempty"`
}
