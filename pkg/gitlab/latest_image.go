package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"golang.org/x/mod/semver"
)

func LatestImage(v []*gitlab.RegistryRepositoryTag) *gitlab.RegistryRepositoryTag {
	result := v[0]

	for _, element := range v {
		current := ImageVersion(element)

		// skip latest
		if current == LatestVersion {
			continue
		}

		if semver.Compare(current, ImageVersion(result)) > 0 {
			result = element
		}
	}

	return result
}
