package manifest

import (
	"fmt"
	"reflect"
	"strings"
)

func (descriptor *Assertion) isValidExpectationValue() error {

	expectedType := strings.ToLower(strings.TrimSpace(descriptor.Expected.Type))

	actualType := strings.ToLower(strings.TrimSpace(reflect.TypeOf(descriptor.Expected.Value).String()))

	if actualType != expectedType {

		return fmt.Errorf("expected type mismatch for descriptor (%s)", descriptor.Label)

	}

	return nil

}
