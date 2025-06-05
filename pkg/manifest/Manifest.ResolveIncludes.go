// file: pkg/manifest/Manifest.ResolveIncludes.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

func (manifest *Manifest) ResolveIncludes() (err error) {

	log.Debugf("ResolveIncludes() loading facts")
	if err = manifest.LoadFacts(); err != nil {
		return err
	}

	log.Debugf("ResolveIncludes() loading patterns")
	if err = manifest.LoadPatterns(); err != nil {
		return err
	}

	log.Debugf("ResolveIncludes() loading assertions")
	if err = manifest.LoadAssertions(); err != nil {
		return err
	}

	return err
}
