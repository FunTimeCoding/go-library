package split

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Space(s string) []string {
	return strings.Split(s, separator.Space)
}
