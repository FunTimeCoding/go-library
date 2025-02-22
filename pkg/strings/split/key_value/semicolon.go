package key_value

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Semicolon(s string) (string, string) {
	result := strings.SplitN(s, separator.Semicolon, 2)
	count := len(result)

	if count == 2 {
		return result[0], result[1]
	} else if count == 1 {
		return result[0], ""
	}

	return "", ""
}
