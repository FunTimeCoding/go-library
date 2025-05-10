package join

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func CommaSpace(s []string) string {
	return strings.Join(s, separator.CommaSpace)
}
