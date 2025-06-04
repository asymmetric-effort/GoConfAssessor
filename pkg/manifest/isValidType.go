// file: pkg/manifest/isValidType.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

func isValidType(t string) bool {
	switch t {
	case "bool", "string", "int", "float", "list", "map":
		return true
	}
	return false
}
