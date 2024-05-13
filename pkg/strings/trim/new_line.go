package trim

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func NewLine(s string) string {
	return strings.TrimRight(s, separator.Unix)
}
