package underscore

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func ToDash(s string) string {
	return strings.ReplaceAll(s, separator.Underscore, separator.Dash)
}
