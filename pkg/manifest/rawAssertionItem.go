// file: pkg/manifest/rawAssertionItem.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// rawAssertionItem - the raw parsed yaml assertion section
type rawAssertionItem struct {
	Label     string                 `yaml:"label"`
	Parser    string                 `yaml:"parser"`
	AppliesTo []string               `yaml:"applies_to"`
	Statement string                 `yaml:"statement"`
	Expected  map[string]interface{} `yaml:"expected"`
	Weight    int                    `yaml:"weight"`
	Source    map[string]string      `yaml:"source"`
	Operator  string                 `yaml:"operator"`
}
