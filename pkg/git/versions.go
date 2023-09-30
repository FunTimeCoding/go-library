package git

import (
	"github.com/coreos/go-semver/semver"
	"strings"
)

func Versions(path string) semver.Versions {
	var result semver.Versions

	for _, element := range Tags(path) {
		if strings.HasPrefix(element, VersionPrefix) {
			result = append(
				result,
				semver.New(strings.TrimPrefix(element, VersionPrefix)),
			)
		}
	}

	semver.Sort(result)

	return result
}
