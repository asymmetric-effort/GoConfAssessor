// file: pkg/manifest/resolve_patterns.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"reflect"
	"strings"
)

// ResolvePatterns replaces any string field whose value begins with "pattern::<name>"
// with the lookup from m.Patterns (using PatternDescriptor.Regex). Errors if undefined
// or if the Regex isn’t a string.
func (manifest *Manifest) ResolvePatterns() error {
	// build name → regex lookup
	patternsMap := make(map[string]interface{}, len(manifest.Patterns))
	for _, pd := range manifest.Patterns {
		patternsMap[pd.Pattern] = pd.Regex
	}
	return replacePatternsInValue(reflect.ValueOf(manifest).Elem(), patternsMap)
}

func replacePatternsInValue(v reflect.Value, patterns map[string]interface{}) error {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			return replacePatternsInValue(v.Elem(), patterns)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if err := replacePatternsInValue(v.Field(i), patterns); err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := replacePatternsInValue(v.Index(i), patterns); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if err := replacePatternsInValue(v.MapIndex(key), patterns); err != nil {
				return err
			}
		}
	case reflect.String:
		s := v.String()
		if strings.HasPrefix(s, "pattern::") {
			name := strings.TrimPrefix(s, "pattern::")
			data, ok := patterns[name]
			if !ok {
				return fmt.Errorf("undefined pattern %q", name)
			}
			regex, ok := data.(string)
			if !ok {
				return fmt.Errorf("pattern %q regex is not a string (got %T)", name, data)
			}
			v.SetString(regex)
		}
	}
	return nil
}
