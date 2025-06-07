// file: pkg/report/Write.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package report

import (
	"errors"
	"github.com/google/uuid"
)

var (
	runId     uuid.UUID
	reportDir string
)

func Start(targetDir string) {
	runId, _ = uuid.NewRandom() // runId is the report prefix
	reportDir = targetDir
}

// Write takes an output path, the general metadata, and the slice of evaluation Results.
// It marshals them into YAML and writes to outPath.
func Write() error {
	// TODO: Transform []evaluator.Result → reportStruct, then marshal to YAML.
	return errors.New("Write: not implemented")
}
