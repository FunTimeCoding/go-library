package dash

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func ToUnderscore(s string) string {
	return strings.ReplaceAll(s, separator.Dash, separator.Underscore)
}
