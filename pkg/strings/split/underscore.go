package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Underscore(s string) []string {
	return strings.Split(s, separator.Underscore)
}
