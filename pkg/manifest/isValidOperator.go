// file: pkg/manifest/isValidType.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

func isValidOperator(op string) bool {
	switch op {
	case "matches", "contains", "equals":
		return true
	}
	return false
}
