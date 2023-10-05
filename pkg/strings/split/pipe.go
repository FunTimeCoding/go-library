package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Pipe(s string) []string {
	return strings.Split(s, separator.Pipe)
}
