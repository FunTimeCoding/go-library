package git

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func ParseProject(path string) (string, string) {
	parts := strings.Split(
		strings.Trim(path, separator.Slash),
		separator.Slash,
	)
	count := len(parts)

	if count != 2 {
		unexpected.Integer(count)
	}

	namespace := parts[0]
	repository := strings.TrimSuffix(parts[count-1], constant.Directory)

	return namespace, repository
}
