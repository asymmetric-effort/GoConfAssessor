// file: pkg/manifest/Manifest.LoadAssertions.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"

	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

// LoadAssertions recursively loads and flattens all assertion-group includes into manifest.Assertions.
// Any entry with Include set is replaced by the contents of that YAML file, which must be a sequence
// of AssertionGroup entries.
func (manifest *Manifest) LoadAssertions() error {
	log := logger.Logger

	// Iterate over manifest.Assertions, expanding includes as we go.
	for i := 0; i < len(manifest.Assertions); i++ {
		group := manifest.Assertions[i]
		if group.Include == "" {
			log.Debugf("AssertionGroup %d has no include", i)
			continue
		}

		log.Debugf("AssertionGroup %d includes '%s'", i, group.Include)
		var nested []AssertionGroup
		if err := utils.LoadYaml(group.Include, &nested); err != nil {
			return fmt.Errorf("failed to load assertions from %s: %w", group.Include, err)
		}

		// Append all loaded groups so they will also be processed for nested includes.
		manifest.Assertions = append(manifest.Assertions, nested...)
	}

	return nil
}
