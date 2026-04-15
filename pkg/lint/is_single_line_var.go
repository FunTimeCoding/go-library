package lint

import "strings"

func isSingleLineVariable(line string) bool {
	if !strings.HasPrefix(line, "var ") {
		return false
	}

	if strings.HasSuffix(line, "{") || strings.HasSuffix(line, "(") {
		return false
	}

	return true
}
