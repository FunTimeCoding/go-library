package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func sourceTypePrefixes(p string) []string {
	parts := strings.Split(strings.Trim(p, separator.Slash), separator.Slash)
	result := make([]string, 0, len(parts))

	for i := len(parts) - 1; i >= 0; i-- {
		result = append(
			result,
			fmt.Sprintf("%s/", strings.Join(parts[:i+1], separator.Slash)),
		)
	}

	result = append(result, "")

	return result
}
