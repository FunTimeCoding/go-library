package maps

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/base64"
)

func ToMarkup(m map[string]string) string {
	result := ""

	for k, v := range m {
		result += fmt.Sprintf("  %s: %s\n", k, base64.Encode(v))
	}

	return result
}
