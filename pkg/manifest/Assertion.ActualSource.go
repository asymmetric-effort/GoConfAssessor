package manifest

type ActualSource struct {
	Path string `yaml:"path,omitempty"`

	Pattern string `yaml:"pattern,omitempty"`
}
