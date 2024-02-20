package release

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/git"
	"strings"
)

func LatestCompatible(
	v []*Release,
	other *semver.Version,
) *Release {
	for _, element := range v {
		s := semver.New(strings.TrimPrefix(element.Name, git.VersionPrefix))

		if s.Major == other.Major && s.Minor == other.Minor {
			return element
		}
	}

	return nil
}
