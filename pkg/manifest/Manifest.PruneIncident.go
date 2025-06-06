package manifest

func (manifest *Manifest) PruneIncludes() (err error) {
	// Prune Facts
	prunedFacts := make([]FactCollection, 0, len(manifest.Facts))
	for _, f := range manifest.Facts {
		if f.Include == "" {
			prunedFacts = append(prunedFacts, f)
		}
	}
	manifest.Facts = prunedFacts

	// Prune Patterns
	prunedPatterns := make([]PatternDescriptor, 0, len(manifest.Patterns))
	for _, p := range manifest.Patterns {
		if p.Include == "" {
			prunedPatterns = append(prunedPatterns, p)
		}
	}
	manifest.Patterns = prunedPatterns

	// Prune Assertions
	prunedAssertions := make([]AssertionGroup, 0, len(manifest.Assertions))
	for _, a := range manifest.Assertions {
		if a.Include == "" {
			prunedAssertions = append(prunedAssertions, a)
		}
	}
	manifest.Assertions = prunedAssertions

	return nil
}
