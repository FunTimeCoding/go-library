package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func SplitSpace(s string) []string {
	return strings.Split(s, separator.Space)
}
