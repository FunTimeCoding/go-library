package store

import (
	"fmt"
	"strings"
)

func sourceTypePrefixes(p string) []string {
	parts := strings.Split(strings.Trim(p, "/"), "/")
	result := make([]string, 0, len(parts))

	for i := len(parts) - 1; i >= 0; i-- {
		result = append(
			result,
			fmt.Sprintf("%s/", strings.Join(parts[:i+1], "/")),
		)
	}

	result = append(result, "")

	return result
}
