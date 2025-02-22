package packages

import (
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
)

func Latest(v []*gitlab.Package) map[string]*gitlab.Package {
	result := make(map[string]*gitlab.Package)

	for _, element := range v {
		latest := result[element.Name]

		if latest == nil ||
			semver.Compare(element.Version, latest.Version) > 0 {
			result[element.Name] = element
		}
	}

	return result
}
