package tag

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
	"strings"
)

func Latest(v []*gitlab.Tag) *gitlab.Tag {
	result := v[0]

	for _, element := range v {
		// only consider tags with prefix
		if !strings.HasPrefix(element.Name, constant.VersionPrefix) {
			continue
		}

		if semver.Compare(element.Name, result.Name) > 0 {
			result = element
		}
	}

	return result
}
