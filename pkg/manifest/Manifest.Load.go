// file: pkg/manifest/Manifest.Load.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/utils"
)

// Load loads the root manifest from manifestFile, resolves includes, prefixes labels, and returns a Manifest.
func (manifest *Manifest) Load(manifestFilename string) (err error) {

	var manifestFile utils.FileEntry
	if manifestFile, err = utils.NewFileEntry(&manifestFilename, true); err != nil {
		return err
	}

	if err = utils.LoadYaml(manifestFile.GetFile(), &manifest); err != nil {
		return fmt.Errorf("failed to unmarshal YAML %s: %w", manifestFile.GetFile(), err)
	}

	if err = manifest.ResolveIncludes(); err != nil {
		return err
	}

	return manifest.Validate()

}
