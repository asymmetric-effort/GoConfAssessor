package manifest

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
	"reflect"
)

// Validate - validate the FactDescriptor
func (descriptor *FactDescriptor) Validate() (err error) {
	if err = utils.ValidIdentifier(descriptor.Fact.Name); err != nil {
		return err
	}

	var allowedTypes = map[reflect.Type]struct{}{
		reflect.TypeOf(""):                  {},
		reflect.TypeOf(int(0)):              {},
		reflect.TypeOf(float64(0)):          {},
		reflect.TypeOf(byte(0)):             {},
		reflect.TypeOf(bool(false)):         {},
		reflect.TypeOf(map[string]string{}): {},
		reflect.TypeOf(map[string]int{}):    {},
		reflect.TypeOf(map[string]bool{}):   {},
		reflect.TypeOf([]string{}):          {},
		reflect.TypeOf([]int{}):             {},
	}
	t := reflect.TypeOf(descriptor.Fact.Data)
	if _, ok := allowedTypes[t]; !ok {
		return fmt.Errorf("invalid fact data type: %v", t)
	}
	v := reflect.ValueOf(descriptor.Fact.Data)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Map:
		if v.Len() == 0 {
			return fmt.Errorf("facts cannot have empty data sets")
		}
	}
	return err
}
