// file: pkg/manifest/Manifest.LoadFacts.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import "github.com/sam-caldwell/GoConfAssessor/pkg/utils"

// LoadFacts - Load facts for the manifest
func (manifest *Manifest) LoadFacts() (err error) {

	var resolve func(theseFacts []FactDescriptor) error

	resolve = func(theseFacts []FactDescriptor) error {
		for _, fact := range theseFacts {
			if childInclude := fact.Include; childInclude != "" {
				var newFacts FactDescriptor
				if err = utils.LoadYaml(childInclude, &newFacts); err != nil {
					return err
				}
				theseFacts = append(theseFacts, newFacts)
				if err := resolve([]FactDescriptor{newFacts}); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return resolve(manifest.Facts)
}
