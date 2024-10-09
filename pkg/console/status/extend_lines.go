package status

import (
	"fmt"
	"strings"
)

func extendLines(
	input string,
	lines []string,
	indent int,
) string {
	result := input

	for _, line := range lines {
		result = fmt.Sprintf(
			"%s\n%s%s\n",
			result,
			spaces(indent),
			line,
		)
	}

	return strings.TrimRight(result, "\n")
}
