// file: pkg/manifest/Assertion.isValidExpectationType.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"strings"
)

func (descriptor *Assertion) isValidExpectationType() error {
	// Added "pattern" as a valid type for regex-based assertions.
	t := strings.ToLower(strings.TrimSpace(descriptor.Expected.Type))
	switch t {
	case "bool", "string", "int", "float", "list", "map", "pattern":
		return nil
	}
	return fmt.Errorf("invalid assertion Expected.Type (%s)", t)
}
