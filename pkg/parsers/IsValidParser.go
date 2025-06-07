package parsers

import (
	"fmt"
	"strings"
)

func IsValidParser(p string) error {
	thisParser := strings.ToLower(strings.TrimSpace(p))
	if _, ok := SupportedParsers[thisParser]; !ok {
		return fmt.Errorf("unknown or unsupported parser (%s)", thisParser)
	}
	return nil
}
