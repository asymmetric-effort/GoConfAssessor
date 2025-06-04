// file: pkg/manifest/RootManifest.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

// RootManifest represents the fully resolved manifest with general metadata and assertion groups.
type RootManifest struct {
	General    map[string]interface{}
	Assertions []Group
}
