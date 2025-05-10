package dash

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func ToUnderscore(s string) string {
	return strings.ReplaceAll(s, separator.Dash, separator.Underscore)
}
