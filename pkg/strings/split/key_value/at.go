package key_value

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func At(s string) (string, string) {
	result := strings.SplitN(s, separator.At, 2)
	count := len(result)

	if count == 2 {
		return result[0], result[1]
	} else if count == 1 {
		return result[0], ""
	}

	return "", ""
}
