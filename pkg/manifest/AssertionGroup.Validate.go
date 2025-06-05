package manifest

import (
	"github.com/sam-caldwell/GoConfAssessor/pkg/parsers"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

// Validate - validate the AssertionGroup
func (descriptor *AssertionGroup) Validate() (err error) {
	if err = utils.ValidIdentifier(descriptor.Name); err != nil {
		return err
	}
	if err = parsers.IsValidParser(descriptor.Parser); err != nil {
		return err
	}
	for _, item := range descriptor.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return err
}
