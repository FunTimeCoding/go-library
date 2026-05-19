package claude

import "strings"

func isSystemNoise(text string) bool {
	trimmed := strings.TrimSpace(text)

	return strings.HasPrefix(trimmed, "<command-name>") ||
		strings.HasPrefix(trimmed, "<local-command") ||
		strings.HasPrefix(trimmed, "<system-reminder>") ||
		strings.Contains(trimmed, "Added") &&
			strings.Contains(trimmed, "as a working directory") ||
		strings.HasPrefix(stripAnsi(trimmed), "Set model to")
}
