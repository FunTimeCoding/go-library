package semver

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"strings"
)

func Trim(version string) string {
	return strings.TrimPrefix(version, constant.VersionPrefix)
}
