package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Slash(s string) []string {
	return strings.Split(s, separator.Slash)
}
