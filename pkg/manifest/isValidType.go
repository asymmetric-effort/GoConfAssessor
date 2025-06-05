// file: pkg/manifest/isValidType.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

func isValidType(t string) bool {
	// Added "pattern" as a valid type for regex-based assertions.
	switch t {
	case "bool", "string", "int", "float", "list", "map", "pattern":
		return true
	}
	return false
}
