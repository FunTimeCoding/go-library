package semver

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"strings"
)

func Trim(version string) string {
	return strings.TrimPrefix(version, git.VersionPrefix)
}
