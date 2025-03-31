package strings

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"regexp"
)

func ReduceSpace(s string) string {
	return regexp.MustCompile(`\s{2,}`).ReplaceAllString(
		s,
		separator.Space,
	)
}
