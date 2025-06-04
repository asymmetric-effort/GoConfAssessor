// file: pkg/evaluator/LoadDeviceConfig.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package evaluator

import "errors"

// LoadDeviceConfig loads and parses a config file at cfgPath using the given parser, returning a DeviceConfig.
func LoadDeviceConfig(cfgPath, parser string) (DeviceConfig, error) {
	// TODO: Implement logic to choose a parser backend (cisco-ios, juniper, yaml, etc.)
	return nil, errors.New("LoadDeviceConfig: not implemented")
}
