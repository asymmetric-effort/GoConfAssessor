// file: pkg/manifest/Manifest.Validate.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

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
	for _, fact := range manifest.Facts {
		if err = fact.Validate(); err != nil {
			return err
		}
	}
	for _, pattern := range manifest.Patterns {
		if err = pattern.Validate(); err != nil {
			return err
		}
	}
	for _, assertion := range manifest.Assertions {
		if err = assertion.Validate(); err != nil {
			return nil
		}
	}

	return nil
}
