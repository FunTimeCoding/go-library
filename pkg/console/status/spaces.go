package status

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func spaces(indent int) string {
	return strings.Repeat(separator.DoubleSpace, indent)
}
