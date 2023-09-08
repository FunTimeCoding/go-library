package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func JoinComma(s []string) string {
	return strings.Join(s, separator.Comma)
}
