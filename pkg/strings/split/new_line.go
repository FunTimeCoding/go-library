package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func NewLine(s string) []string {
	return strings.Split(s, separator.Unix)
}
