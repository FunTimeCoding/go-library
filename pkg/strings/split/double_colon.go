package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func DoubleColon(s string) []string {
	return strings.Split(s, separator.DoubleColon)
}
