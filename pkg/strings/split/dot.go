package split

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Dot(s string) []string {
	return strings.Split(s, separator.Dot)
}
