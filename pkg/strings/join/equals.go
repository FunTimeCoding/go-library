package join

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Equals(s []string) string {
	return strings.Join(s, separator.Equals)
}
