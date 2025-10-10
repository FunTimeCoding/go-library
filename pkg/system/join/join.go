package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/slice"
	"path"
	"strings"
)

func Join(parts ...string) string {
	result := slice.StripEmpty(parts)
	absolute := strings.HasPrefix(result[0], separator.Slash)
	result = slice.Trim(result, separator.Slash)

	if len(result) == 0 {
		return separator.Slash
	}

	if absolute {
		result = slice.Prepend(result, separator.Slash)
	}

	return path.Join(result...)
}
