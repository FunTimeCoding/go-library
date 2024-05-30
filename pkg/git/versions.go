package git

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"strings"
)

func Versions(path string) semver.Versions {
	var result semver.Versions

	for _, element := range Tags(path) {
		if strings.HasPrefix(element, constant.VersionPrefix) {
			result = append(
				result,
				semver.New(
					strings.TrimPrefix(
						element, constant.VersionPrefix,
					),
				),
			)
		}
	}

	semver.Sort(result)

	return result
}
