// file: pkg/report/Write.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package report

import "errors"

// Write takes an output path, the general metadata, and the slice of evaluation Results.
// It marshals them into YAML and writes to outPath.
func Write(outPath string, general map[string]interface{}, results []evaluator.Result) error {
	// TODO: Transform []evaluator.Result → reportStruct, then marshal to YAML.
	return errors.New("Write: not implemented")
}
