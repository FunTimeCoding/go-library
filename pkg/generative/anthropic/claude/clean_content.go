package claude

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"strings"
)

func cleanContent(s string) string {
	s = ansiPattern.ReplaceAllString(s, "")
	s = markupTagPattern.ReplaceAllString(s, " ")
	s = join.Space(strings.Fields(s)...)

	return strings.TrimSpace(s)
}
