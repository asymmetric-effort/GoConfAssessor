// file: pkg/manifest/resolve_facts.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"reflect"
	"strings"
)

func (manifest *Manifest) ResolveFacts() error {
	// build lookup: fact name → its Data
	factsMap := make(map[string]interface{}, len(manifest.Facts))
	for _, fc := range manifest.Facts {
		factsMap[fc.Fact] = fc.Data
	}
	log.Debug("Starting to replace facts in value fields")
	return replaceFactsInValue(reflect.ValueOf(manifest).Elem(), factsMap)
}

func replaceFactsInValue(v reflect.Value, facts map[string]interface{}) error {
	switch v.Kind() {

	// — if this field is an interface, see if it wraps a "fact::…" string.
	//   if so, assign the raw Data (whatever its Go type) right back into the interface.
	case reflect.Interface:
		if v.CanSet() && !v.IsNil() {
			elem := v.Elem()
			if elem.Kind() == reflect.String {
				s := elem.String()
				if strings.HasPrefix(s, "fact::") {
					name := strings.TrimPrefix(s, "fact::")
					data, ok := facts[name]
					if !ok {
						return fmt.Errorf("undefined fact %q", name)
					}
					v.Set(reflect.ValueOf(data)) // <— handles []interface{}, map[string]interface{}, bool, etc.
					return nil
				}
			}
			// otherwise, dig into whatever the interface currently holds
			return replaceFactsInValue(elem, facts)
		}

	case reflect.Ptr:
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
			if err := replaceFactsInValue(v.MapIndex(key), facts); err != nil {
				return err
			}
		}

	// — for plain string fields (e.g. your struct fields typed `string`),
	//   continue to only support replacing with string facts.
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
