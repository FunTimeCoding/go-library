package github

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/google/go-github/v69/github"
	"golang.org/x/mod/semver"
	"strings"
)

func LatestTag(v []*github.RepositoryTag) *github.RepositoryTag {
	var tags []string

	for _, element := range v {
		// Only consider tags with prefix
		if strings.HasPrefix(*element.Name, constant.VersionPrefix) {
			tags = append(tags, *element.Name)
		}
	}

	semver.Sort(tags)
	count := len(tags)
	var latest *github.RepositoryTag
	index := count - 1

	for _, element := range v {
		if *element.Name == tags[index] {
			latest = element
		}
	}

	return latest
}
