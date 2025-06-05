package manifest

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

// Validate - validate the PatternDescriptor
func (descriptor *PatternDescriptor) Validate() (err error) {
	if err = utils.ValidIdentifier(descriptor.Pattern); err != nil {
		return err
	}
	if descriptor.Regex == "" {
		return fmt.Errorf("pattern regex cannot be empty")
	}
	return err
}
