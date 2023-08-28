package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func SplitNewLine(s string) []string {
	return strings.Split(s, separator.Unix)
}
