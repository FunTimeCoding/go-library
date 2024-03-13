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
	index := len(versions) - 1
	var result *gitlab.RegistryRepositoryTag

	for _, element := range v {
		if strings.HasSuffix(element.Path, versions[index]) {
			result = element
		}
	}

	return result
}
