package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func SplitComma(s string) []string {
	return strings.Split(s, separator.Comma)
}
