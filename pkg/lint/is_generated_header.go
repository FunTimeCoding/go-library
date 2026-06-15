package lint

import "strings"

func IsGeneratedHeader(content string) bool {
	for _, line := range strings.SplitN(content, "\n", 10) {
		trimmed := strings.TrimSpace(line)

		if !strings.HasPrefix(trimmed, "//") && trimmed != "" {
			return false
		}

		if strings.HasPrefix(trimmed, "// Code generated") &&
			strings.HasSuffix(trimmed, "DO NOT EDIT.") {
			return true
		}
	}

	return false
}
