package packages

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
)

func Latest(v []*gitlab.Package) map[string]*gitlab.Package {
	result := make(map[string]*gitlab.Package)

	for _, e := range v {
		errors.PanicOnSemver(e.Version)

		latest := result[e.Name]

		if latest == nil ||
			semver.Compare(e.Version, latest.Version) > 0 {
			result[e.Name] = e
		}
	}

	return result
}
