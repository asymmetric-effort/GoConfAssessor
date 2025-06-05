// file: pkg/manifest/GeneralSection.Validate.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"errors"
	"fmt"
	"strings"
)

// Validate - Validate the general section
func (g *GeneralSection) Validate() (err error) {

	// General.name must be a non-empty string
	if strings.TrimSpace(g.Name) == "" {
		return errors.New("manifest general.name must be a non-empty string")
	}

	// General.version must follow semantic versioning (e.g., 1.0.0)
	version := g.Version
	if !isValidSemVer(version) {
		return fmt.Errorf("manifest general.version invalid semver: %v", version)
	}

	// Metadata entries, if present, must be key->string
	if err = g.validateMetadata(); err != nil {
		return err
	}

	return err
}
