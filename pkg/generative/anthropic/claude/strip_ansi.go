package claude

func stripAnsi(s string) string {
	return ansiPattern.ReplaceAllString(s, "")
}
