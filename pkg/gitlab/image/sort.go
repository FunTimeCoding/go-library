package image

import (
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
	"sort"
)

func Sort(v []*gitlab.RegistryRepositoryTag) []*gitlab.RegistryRepositoryTag {
	result := append([]*gitlab.RegistryRepositoryTag{}, v...)

	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return semver.Compare(Version(result[i]), Version(result[j])) > 0
		},
	)

	return result
}
