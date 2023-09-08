package split

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Space(s string) []string {
	return strings.Split(s, separator.Space)
}
