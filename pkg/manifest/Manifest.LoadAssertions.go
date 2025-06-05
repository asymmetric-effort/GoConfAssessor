// file: pkg/manifest/Manifest.LoadAssertions.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import "github.com/sam-caldwell/GoConfAssessor/pkg/utils"

func (manifest *Manifest) LoadAssertions() (err error) {
	var resolve func(theseAssertions []AssertionGroup) error
	resolve = func(theseAssertions []AssertionGroup) error {
		for _, group := range theseAssertions {
			if childInclude := group.Include; childInclude != "" {
				var newAssertionGroup AssertionGroup
				if err = utils.LoadYaml(childInclude, &newAssertionGroup); err != nil {
					return err
				}
				theseAssertions = append(theseAssertions, newAssertionGroup)
				if err := resolve([]AssertionGroup{newAssertionGroup}); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return resolve(manifest.Assertions)
}
