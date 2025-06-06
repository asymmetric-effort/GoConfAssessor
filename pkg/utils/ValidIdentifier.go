package utils

import (
	"fmt"
	"regexp"
)

// ValidIdentifier - validating the identifier pattern
func ValidIdentifier(s string) error {

	const identifierPattern = "^[a-zA-Z]{0,1}[a-zA-Z0-9_\\-\\.]{0,63}[a-zA-Z0-9]$"

	if pattern := regexp.MustCompile(identifierPattern); pattern.MatchString(s) {

		return nil

	}

	return fmt.Errorf("identifier (%s) invalid.  Must match %s", s, identifierPattern)

}
