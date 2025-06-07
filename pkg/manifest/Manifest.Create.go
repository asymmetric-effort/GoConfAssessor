// file: pkg/manifest/Manifest.Create.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import (
	"fmt"
	"time"
)

// Create loads the root manifest from manifestFile, resolves includes, prefixes labels, and returns a Manifest.
func (manifest *Manifest) Create(manifestFilename, name, author *string) (err error) {
	log.Debug("Creating %s", manifestFilename)
	manifest.General = GeneralSection{
		Name:    *name,
		Version: Version,
		Metadata: map[string]string{
			"author":     *author,
			"created_on": fmt.Sprintf("%s", time.Now()),
			"created_by": "Manifest.Create",
		},
	}
	manifest.Facts = []FactCollection{
		{
			Fact: "example1",
			Data: "value1",
		},
		{
			Fact: "example2",
			Data: "value2",
		},
		{
			Fact: "example3",
			Data: "value3",
		},
	}
	manifest.Patterns = []PatternDescriptor{
		{Pattern: "examplePattern",
			Regex: ".+",
		},
	}
	manifest.Assertions = []AssertionGroup{
		{
			Name:   "example_assertion",
			Parser: "text",
			Items: []Assertion{
				{
					Label:     "example_check",
					Statement: "A check is an expected measurable state against which an actual state can be compared",
					Expected: Expectation{
						Type:  "string",
						Value: "expected_value",
					},
					Weight: 1,
					Source: ActualSource{
						Path:    "dot-delimited-path-to-a-config-object from the source file",
						Pattern: "a regular expression to find some pattern in the source file",
					},
				},
			},
		},
	}
	return nil
}
