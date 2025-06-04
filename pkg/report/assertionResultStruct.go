// file: pkg/report/assertionResultStruct.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package report

type assertionResultStruct struct {
	Group     string      `yaml:"group"`
	Label     string      `yaml:"label"`
	Statement string      `yaml:"statement"`
	Expected  interface{} `yaml:"expected"`
	Actual    interface{} `yaml:"actual"`
	Weight    int         `yaml:"weight"`
	Passed    bool        `yaml:"passed"`
}
