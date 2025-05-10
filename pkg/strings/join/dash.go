package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Dash(s []string) string {
	return strings.Join(s, separator.Dash)
}
