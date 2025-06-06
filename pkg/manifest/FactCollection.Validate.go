// file: pkg/manifest/FactCollection.Validate.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

// Validate checks that the fact name is valid and that Fact.Data is one of the supported types.
// It also ensures non‐empty data for slices, maps, and strings.
func (descriptor *FactCollection) Validate() error {
	if descriptor.Include != "" {
		return nil
	}
	if err := utils.ValidIdentifier(descriptor.Fact); err != nil {
		return err
	}

	switch v := descriptor.Data.(type) {
	case string:
		if v == "" {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case int:
		if v == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case float64:
		if v == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case byte:
		if v == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case bool:
		// both true and false are allowed
		return nil

	case map[string]interface{}:
		if len(v) == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case []string:
		if len(v) == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case []int:
		if len(v) == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		return nil

	case []interface{}:
		if len(v) == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
		for idx, elem := range v {
			switch elem.(type) {
			case string, int, float64, byte, bool:
				// OK
			default:
				return fmt.Errorf(
					"invalid element type %T in slice for fact %q at index %d",
					elem, descriptor.Fact, idx,
				)
			}
		}
		return nil

	default:
		return fmt.Errorf("invalid fact data type: %T", v)
	}
}
