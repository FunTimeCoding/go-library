package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Comma(s string) []string {
	return strings.Split(s, separator.Comma)
}
