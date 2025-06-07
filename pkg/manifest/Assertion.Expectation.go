package manifest

type Expectation struct {
	Type string `yaml:"type,omitempty"`

	Value interface{} `yaml:"value,omitempty"`
}
