package lint

import (
	"fmt"
	"strings"
)

func toSingleLine(lines []string) string {
	if len(lines) == 0 {
		return ""
	}

	first := lines[0]
	open := strings.Index(first, "{")

	if open == -1 {
		return first
	}

	return fmt.Sprintf("%s {}", strings.TrimSpace(first[:open]))
}
