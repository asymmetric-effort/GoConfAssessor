// file: pkg/evaluator/RunAll.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package evaluator

import "errors"

// RunAll evaluates every assertion in assertionGroups against deviceCfg and returns a slice of Result.
func RunAll(assertionGroups []manifest.Group, deviceCfg DeviceConfig) ([]Result, error) {
	// TODO: Iterate through each group and each assertion, call Evaluate, and collect results.
	return nil, errors.New("RunAll: not implemented")
}
