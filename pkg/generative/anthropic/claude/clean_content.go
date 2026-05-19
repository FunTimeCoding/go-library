package claude

import "strings"

func cleanContent(s string) string {
	s = ansiPattern.ReplaceAllString(s, "")
	s = markupTagPattern.ReplaceAllString(s, " ")
	s = strings.Join(strings.Fields(s), " ")

	return strings.TrimSpace(s)
}
