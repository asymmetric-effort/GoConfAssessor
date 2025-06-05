// file: pkg/manifest/Manifest.Validate.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"errors"
	"fmt"
)

// Validate checks that the Manifest contains valid fields and formats.
// This program will verify each field of rootManifest for legitimate values:
//   - general.name non-empty
//   - general.version follows semantic version format
//   - metadata entries are key->string
//   - each group has non-empty Name
//   - each assertion:
//   - Label non-empty
//   - Parser is one of the allowed parsers
//   - AppliesTo non-empty
//   - Statement non-empty
//   - Expected.Type is valid and Expected.Value matches type
//   - Weight >= 0
//   - Source has at least Path or Pattern
//   - Operator is one of matches|contains|equals
func (manifest *Manifest) Validate() (err error) {

	// validate the general section
	if err = manifest.General.Validate(); err != nil {
		return err
	}

	// Validate each group and its assertions
	for _, group := range manifest.Assertions {
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
