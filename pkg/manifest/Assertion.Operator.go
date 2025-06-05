package manifest

import (
	"fmt"
	"strings"
)

func (descriptor *Assertion) isValidOperator() (err error) {
	switch o := strings.ToLower(strings.TrimSpace(descriptor.Operator)); o {
	case "contains", "matches", "not_matches", "excludes", "equals", "not_equals":
		return nil
	default:
		return fmt.Errorf("invalid assertion operator (%s). see assertion (%s)", o, descriptor.Label)
	}
}
