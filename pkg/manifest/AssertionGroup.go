// file: pkg/manifest/Group.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// Group represents a named collection of assertions.
type Group struct {
	Name  string
	Items []Assertion
}
