package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func JoinNewLine(s []string) string {
	return strings.Join(s, separator.Unix)
}
