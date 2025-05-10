package compact

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"strings"
)

func Mail(s string) string {
	if strings.Contains(s, separator.At) {
		first, _ := key_value.At(s)

		return first
	}

	return s
}
