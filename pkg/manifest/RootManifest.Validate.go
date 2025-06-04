// file: pkg/manifest/RootManifest.Validate.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"errors"
	"fmt"
)

// Validate checks that the RootManifest contains valid fields and formats.
func (rm *RootManifest) Validate() error {
	// General.name must be a non-empty string
	name, ok := rm.General["name"].(string)
	if !ok || name == "" {
		return errors.New("manifest general.name must be a non-empty string")
	}

	// General.version must follow semantic versioning (e.g., 1.0.0)
	version, ok := rm.General["version"].(string)
	if !ok || !isValidSemVer(version) {
		return fmt.Errorf("manifest general.version invalid semver: %v", rm.General["version"])
	}

	// Metadata entries, if present, must be key->string
	if metaList, exists := rm.General["metadata"]; exists {
		entries, ok := metaList.([]interface{})
		if !ok {
			return errors.New("manifest general.metadata must be a sequence of key->string maps")
		}
		for _, entry := range entries {
			m, ok := entry.(map[string]interface{})
			if !ok || len(m) != 1 {
				return errors.New("each metadata entry must be a single key->string map")
			}
			for _, v := range m {
				if _, ok := v.(string); !ok {
					return errors.New("metadata values must be strings")
				}
			}
		}
	}

	// Validate each group and its assertions
	for _, group := range rm.Assertions {
		if group.Name == "" {
			return errors.New("each group must have a non-empty Name")
		}
		for _, a := range group.Items {
			if a.Label == "" {
				return fmt.Errorf("assertion in group %q has empty Label", group.Name)
			}
			// Parser must be one of the allowed set
			if !isValidParser(a.Parser) {
				return fmt.Errorf("assertion %q: invalid parser %q", a.Label, a.Parser)
			}
			if len(a.AppliesTo) == 0 {
				return fmt.Errorf("assertion %q: AppliesTo must be non-empty", a.Label)
			}
			if a.Statement == "" {
				return fmt.Errorf("assertion %q: Statement must be non-empty", a.Label)
			}
			// Expected.Type must be valid and Value non-nil
			t := a.Expected.Type
			if !isValidType(t) || a.Expected.Value == nil {
				return fmt.Errorf("assertion %q: Expected.Type invalid or Value missing", a.Label)
			}
			if a.Weight < 0 {
				return fmt.Errorf("assertion %q: Weight must be >= 0", a.Label)
			}
			// Source must have at least Path or Pattern
			if a.Source.Path == "" && a.Source.Pattern == "" {
				return fmt.Errorf("assertion %q: Source must have a Path or Pattern", a.Label)
			}
			// Operator must be one of matches, contains, equals
			if !isValidOperator(a.Operator) {
				return fmt.Errorf("assertion %q: invalid Operator %q", a.Label, a.Operator)
			}
		}

	}
	return nil
}
