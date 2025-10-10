package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/slice"
	"path"
)

func Relative(parts ...string) string {
	return path.Join(slice.Trim(slice.StripEmpty(parts), separator.Slash)...)
}
