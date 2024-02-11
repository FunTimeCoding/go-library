package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/mod/semver"
	"strings"
)

func LatestTag(v []*gitlab.Tag) *gitlab.Tag {
	var tags []string

	for _, element := range v {
		// Only consider tags with prefix
		if strings.HasPrefix(element.Name, git.VersionPrefix) {
			tags = append(tags, element.Name)
		}
	}

	semver.Sort(tags)
	count := len(tags)
	var latest *gitlab.Tag
	index := count - 1

	for _, element := range v {
		if element.Name == tags[index] {
			latest = element
		}
	}

	return latest
}
