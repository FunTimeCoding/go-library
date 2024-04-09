package image

import (
	library "github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/mod/semver"
)

func Latest(v []*gitlab.RegistryRepositoryTag) *gitlab.RegistryRepositoryTag {
	result := v[0]

	for _, element := range v {
		current := Version(element)

		// skip latest
		if current == library.LatestVersion {
			continue
		}

		if semver.Compare(current, Version(result)) > 0 {
			result = element
		}
	}

	return result
}
