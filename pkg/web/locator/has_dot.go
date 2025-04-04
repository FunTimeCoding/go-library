package locator

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func HasDot(s string) bool {
	return strings.Count(s, separator.Dot) > 1
}
