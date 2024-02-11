package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"golang.org/x/mod/semver"
	"strings"
)

func LatestImage(v []*gitlab.RegistryRepositoryTag) *gitlab.RegistryRepositoryTag {
	var versions []string

	for _, element := range v {
		parts := strings.Split(element.Path, ":")
		version := parts[1]

		if version == "latest" {
			continue
		}

		versions = append(versions, version)
	}

	semver.Sort(versions)
	count := len(versions)
	var latest *gitlab.RegistryRepositoryTag
	index := count - 1

	for _, element := range v {
		if strings.HasSuffix(element.Path, versions[index]) {
			latest = element
		}
	}

	return latest
}
