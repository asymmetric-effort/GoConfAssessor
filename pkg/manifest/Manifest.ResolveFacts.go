package manifest

import (
	"fmt"
	"reflect"
	"strings"
)

func (manifest *Manifest) ResolveFacts() error {
	// build a quick lookup: fact name → its Data
	factsMap := make(map[string]interface{}, len(manifest.Facts))
	for _, fc := range manifest.Facts {
		factsMap[fc.Fact] = fc.Data
	}

	// start recursive replacement
	return replaceFactsInValue(reflect.ValueOf(manifest).Elem(), factsMap)
}

func replaceFactsInValue(v reflect.Value, facts map[string]interface{}) error {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			return replaceFactsInValue(v.Elem(), facts)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if err := replaceFactsInValue(v.Field(i), facts); err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := replaceFactsInValue(v.Index(i), facts); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			if err := replaceFactsInValue(val, facts); err != nil {
				return err
			}
		}
	case reflect.String:
		s := v.String()
		if strings.HasPrefix(s, "fact::") {
			name := strings.TrimPrefix(s, "fact::")
			data, ok := facts[name]
			if !ok {
				return fmt.Errorf("undefined fact %q", name)
			}
			str, ok := data.(string)
			if !ok {
				return fmt.Errorf("fact %q data is not a string (got %T)", name, data)
			}
			v.SetString(str)
		}
	}
	return nil
}
