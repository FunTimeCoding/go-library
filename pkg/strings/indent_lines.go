package strings

import (
	"fmt"
	"strings"
)

func IndentLines(
	input []string,
	count int,
) []string {
	var result []string
	indent := strings.Repeat(" ", count)

	for _, element := range input {
		result = append(result, fmt.Sprintf("%s%s", indent, element))
	}

	return result
}
