package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Slash(s []string) string {
	return strings.Join(s, separator.Slash)
}
