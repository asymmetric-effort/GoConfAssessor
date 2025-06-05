package manifest

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
	"strings"
)

func (descriptor *Assertion) Validate() (err error) {
	if err = utils.ValidIdentifier(descriptor.Label); err != nil {
		return err
	}
	if statement := strings.TrimSpace(descriptor.Statement); statement == "" {
		return fmt.Errorf("assertions cannot have missing or empty statement fields")
	}
	if err = descriptor.isValidExpectationType(); err != nil {
		return err
	}
	if err = descriptor.isValidExpectationValue(); err != nil {
		return err
	}
	if err = descriptor.validateSource(); err != nil {
		return err
	}
	if err = descriptor.isValidOperator(); err != nil {
		return err
	}
	return err
}
