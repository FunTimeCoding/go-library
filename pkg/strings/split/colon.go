package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Colon(s string) []string {
	return strings.Split(s, separator.Colon)
}
