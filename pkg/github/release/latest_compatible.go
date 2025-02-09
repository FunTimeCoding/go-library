package release

import (
	"github.com/coreos/go-semver/semver"
	library "github.com/funtimecoding/go-library/pkg/semver"
)

func LatestCompatible(
	v []*Release,
	other *semver.Version,
) *Release {
	for _, element := range v {
		s := semver.New(library.Trim(element.Name))

		if s.Major == other.Major && s.Minor == other.Minor {
			return element
		}
	}

	return nil
}
