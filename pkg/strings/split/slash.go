package split

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Slash(s string) []string {
	return strings.Split(s, separator.Slash)
}
