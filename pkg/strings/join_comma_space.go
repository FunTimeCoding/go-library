package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func JoinCommaSpace(s []string) string {
	return strings.Join(s, separator.CommaSpace)
}
