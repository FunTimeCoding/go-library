package image

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
)

func Latest(v []*gitlab.RegistryRepositoryTag) *gitlab.RegistryRepositoryTag {
	result := v[0]

	for _, e := range v {
		current := Version(e)

		// skip latest
		if current == constant.LatestVersion {
			continue
		}

		if semver.Compare(current, Version(result)) > 0 {
			result = e
		}
	}

	return result
}
