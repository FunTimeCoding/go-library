package join

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func NewLine(s []string) string {
	return strings.Join(s, separator.Unix)
}
