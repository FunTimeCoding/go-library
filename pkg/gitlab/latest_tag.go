package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/mod/semver"
	"strings"
)

func LatestTag(v []*gitlab.Tag) *gitlab.Tag {
	result := v[0]

	for _, element := range v {
		// only consider tags with prefix
		if !strings.HasPrefix(element.Name, git.VersionPrefix) {
			continue
		}

		if semver.Compare(element.Name, result.Name) > 0 {
			result = element
		}
	}

	return result
}
