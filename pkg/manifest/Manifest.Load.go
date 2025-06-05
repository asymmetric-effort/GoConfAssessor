// file: pkg/manifest/LoadAndResolve.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

// LoadAndResolve loads the root manifest from manifestFile, resolves includes, prefixes labels, and returns a RootManifest.
func LoadAndResolve(manifestFile string) (RootManifest, error) {
	absPath, err := filepath.Abs(manifestFile)
	if err != nil {
		return RootManifest{}, fmt.Errorf("failed to get absolute path: %w", err)
	}
	dir := filepath.Dir(absPath)

	// fileEntry is used to track each file to process.
	type fileEntry struct {
		path    string
		baseDir string
		prefix  string
		isRoot  bool
	}

	filesCh := make(chan fileEntry, 100)
	resultsCh := make(chan []Group, 100)

	var general GeneralSection
	var allGroups []Group
	initialized := false

	// Enqueue the root manifest
	filesCh <- fileEntry{path: absPath, baseDir: dir, prefix: "", isRoot: true}

	// Process until no more files
	for len(filesCh) > 0 {
		fe := <-filesCh

		data, err := os.ReadFile(fe.path)
		if err != nil {
			return RootManifest{}, fmt.Errorf("failed to read file %s: %w", fe.path, err)
		}

		var raw rawManifest
		if err := yaml.Unmarshal(data, &raw); err != nil {
			return RootManifest{}, fmt.Errorf("failed to unmarshal YAML %s: %w", fe.path, err)
		}

		if fe.isRoot {
			general = raw.General
			initialized = true
		}

		var batch []Group
		for _, entry := range raw.Assertions {
			if entry.Include != "" {
				includePath := entry.Include
				if !filepath.IsAbs(includePath) {
					includePath = filepath.Join(fe.baseDir, includePath)
				}
				includeAbs, err := filepath.Abs(includePath)
				if err != nil {
					return RootManifest{}, fmt.Errorf("failed to get abs path of include %s: %w", includePath, err)
				}
				includeDir := filepath.Dir(includeAbs)
				base := filepath.Base(includeAbs)
				baseName := strings.TrimSuffix(base, filepath.Ext(base))
				newPrefix := fe.prefix + baseName + "/"

				filesCh <- fileEntry{path: includeAbs, baseDir: includeDir, prefix: newPrefix, isRoot: false}
			} else {
				var items []Assertion
				for _, rawItem := range entry.Items {
					var a Assertion
					a.Label = fe.prefix + rawItem.Label
					a.Parser = rawItem.Parser
					a.AppliesTo = rawItem.AppliesTo
					a.Statement = rawItem.Statement
					if t, ok := rawItem.Expected["type"].(string); ok {
						a.Expected.Type = t
					}
					if v, ok := rawItem.Expected["value"]; ok {
						a.Expected.Value = v
					}
					// Assign weight directly
					a.Weight = rawItem.Weight
					if p, ok := rawItem.Source["path"]; ok {
						a.Source.Path = p
					}
					if pat, ok := rawItem.Source["pattern"]; ok {
						a.Source.Pattern = pat
					}
					a.Operator = rawItem.Operator
					items = append(items, a)
				}
				batch = append(batch, Group{Name: entry.Name, Items: items})
			}
		}
		resultsCh <- batch
	}

	close(resultsCh)
	for groups := range resultsCh {
		allGroups = append(allGroups, groups...)
	}

	if !initialized {
		return RootManifest{}, fmt.Errorf("no general section found in root manifest %s", manifestFile)
	}
	return RootManifest{General: general, Assertions: allGroups}, nil
}
