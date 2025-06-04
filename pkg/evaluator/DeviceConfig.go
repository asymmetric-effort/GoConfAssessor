// file: pkg/evaluator/DeviceConfig.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package evaluator

// DeviceConfig abstracts a device's parsed configuration.
type DeviceConfig interface {
	// Lookup returns the value found at the given path or an error.
	Lookup(path string) (interface{}, error)
	// RawContent returns the raw configuration text for regex-based parses.
	RawContent() (string, error)
}
