// file: pkg/manifest/Manifest.LoadPatterns.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

func (m *Manifest) LoadPatterns() error {
	log := logger.Logger

	for i := 0; i < len(m.Patterns); i++ {
		p := m.Patterns[i]
		if p.Include == "" {
			log.Debugf("Pattern %d has no include", i)
			continue
		}

		log.Debugf("Pattern %d includes '%s'", i, p.Include)
		var nested []PatternDescriptor
		if err := utils.LoadYaml(p.Include, &nested); err != nil {
			return fmt.Errorf("failed to load patterns from %s: %w", p.Include, err)
		}

		// Append all loaded patterns so they also get checked for nested includes (if you ever support them)
		m.Patterns = append(m.Patterns, nested...)
	}

	return nil
}
