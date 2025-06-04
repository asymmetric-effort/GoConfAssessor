// file: pkg/manifest/isValidSemVer.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package manifest

import "regexp"

func isValidSemVer(v string) bool {
	// Basic regex for semver: major.minor.patch (no pre-release/build)
	re := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)
	return re.MatchString(v)
}
