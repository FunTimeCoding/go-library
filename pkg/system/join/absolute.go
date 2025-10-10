package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/slice"
	"path"
	"strings"
)

func Absolute(parts ...string) string {
	parts = slice.Trim(slice.StripEmpty(parts), separator.Slash)

	if len(parts) == 0 {
		return separator.Slash
	}

	result := path.Join(parts...)

	if !strings.HasPrefix(result, separator.Slash) {
		result = key_value.Empty(separator.Slash, result)
	}

	return result
}
