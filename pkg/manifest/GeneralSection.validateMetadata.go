// file: pkg/manifest/GeneralSection.GeneralSection.validateMetadata.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// validateMetadata - validate the general metadata section
func (g *GeneralSection) validateMetadata() (err error) {

	if g.Metadata == nil {
		// initialize to avoid future errors
		g.Metadata = make(map[string]string)
		return nil
	}

	const keyPattern = `^[A-Za-z][A-Za-z0-9_]{1,63}$`
	re := regexp.MustCompile(keyPattern)

	for key, value := range g.Metadata {
		// Key must not be empty or only whitespace
		if strings.TrimSpace(key) == "" {
			return errors.New("metadata keys cannot be whitespace or empty")
		}
		// Key must match the allowed pattern
		if !re.MatchString(key) {
			return fmt.Errorf("metadata key %q must match pattern %s", key, keyPattern)
		}
		// Value must not be empty
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("metadata value for key %q cannot be empty", key)
		}
		// Enforce maximum length
		if len(value) > 1024 {
			return fmt.Errorf("metadata value for key %q exceeds 1024 bytes", key)
		}
	}

	return nil
}
