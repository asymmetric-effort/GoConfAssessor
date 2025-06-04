// file: pkg/report/report.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package report

// reportStruct mirrors the final report YAML format.
type reportStruct struct {
	General    map[string]interface{}  `yaml:"general"`
	Assertions []assertionResultStruct `yaml:"assertions"`
}
