package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Colon(s ...string) string {
	return strings.Join(s, separator.Colon)
}
