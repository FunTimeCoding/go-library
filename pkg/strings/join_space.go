package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func JoinSpace(s []string) string {
	return strings.Join(s, separator.Space)
}
