// file: pkg/manifest/Manifest.LoadPatterns.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import "github.com/sam-caldwell/GoConfAssessor/pkg/utils"

func (manifest *Manifest) LoadPatterns() (err error) {

	var resolve func(thesePatterns []PatternDescriptor) error

	resolve = func(thesePatterns []PatternDescriptor) error {

		for _, fact := range thesePatterns {

			if childInclude := fact.Include; childInclude != "" {

				var newPatterns PatternDescriptor
				if err = utils.LoadYaml(childInclude, &newPatterns); err != nil {
					return err
				}

				thesePatterns = append(thesePatterns, newPatterns)
				if err := resolve([]PatternDescriptor{newPatterns}); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return resolve(manifest.Patterns)
}
