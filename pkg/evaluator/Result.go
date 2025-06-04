// file: pkg/evaluator/Result.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package evaluator

// Result holds one assertion’s evaluation outcome.
type Result struct {
	Group     string
	Label     string
	Statement string
	Expected  interface{}
	Actual    interface{}
	Weight    int
	Passed    bool
	Error     error
}
