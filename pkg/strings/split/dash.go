package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Dash(s string) []string {
	return strings.Split(s, separator.Dash)
}
