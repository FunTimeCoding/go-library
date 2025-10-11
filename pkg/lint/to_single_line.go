package lint

import "strings"

func toSingleLine(lines []string) string {
	if len(lines) == 0 {
		return ""
	}

	first := lines[0]
	open := strings.Index(first, "{")

	if open == -1 {
		return first
	}

	return strings.TrimSpace(first[:open]) + " {}"
}
