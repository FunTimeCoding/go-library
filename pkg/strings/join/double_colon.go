package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func DoubleColon(s []string) string {
	return strings.Join(s, separator.DoubleColon)
}
