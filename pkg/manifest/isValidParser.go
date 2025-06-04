// file: pkg/manifest/isValidParser.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

func isValidParser(p string) bool {
	switch p {
	case "cisco-ios", "juniper", "arista", "yaml", "json", "ini", "text":
		return true
	}
	return false
}
