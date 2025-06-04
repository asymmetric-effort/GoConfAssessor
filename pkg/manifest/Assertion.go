// file: pkg/manifest/Assertion.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// Assertion represents a single check in a group.
type Assertion struct {
	Label     string
	Parser    string
	AppliesTo []string
	Statement string
	Expected  struct {
		Type  string
		Value interface{}
	}
	Weight int
	Source struct {
		Path    string
		Pattern string
	}
	Operator string
}
