package strings

import (
	"fmt"
	"strings"
)

func IndentLines(
	v []string,
	count int,
) []string {
	var result []string
	indent := strings.Repeat(" ", count)

	for _, e := range v {
		result = append(result, fmt.Sprintf("%s%s", indent, e))
	}

	return result
}
