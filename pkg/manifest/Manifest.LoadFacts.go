// file: pkg/manifest/Manifest.LoadFacts.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"

	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

// LoadFacts recursively loads and flattens all fact includes into manifest.Facts.
// Any entry with Include set is replaced by the contents of that YAML file.
// The included file must contain a YAML sequence of “- fact:” entries.
func (manifest *Manifest) LoadFacts() error {
	log := logger.Logger

	// Iterate over manifest.Facts, expanding includes as we go.
	for i := 0; i < len(manifest.Facts); i++ {
		f := manifest.Facts[i]
		if f.Include == "" {
			log.Debugf("Fact %d has no include", i)
			continue
		}

		log.Debugf("Fact %d includes '%s'", i, f.Include)
		var nested []FactDescriptor
		if err := utils.LoadYaml(f.Include, &nested); err != nil {
			return fmt.Errorf("failed to load facts from %s: %w", f.Include, err)
		}

		// Append all loaded facts so they will also be processed for nested includes.
		manifest.Facts = append(manifest.Facts, nested...)
	}

	return nil
}
