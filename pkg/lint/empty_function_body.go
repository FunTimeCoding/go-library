package lint

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func emptyFunctionBody(lines []string) bool {
	if len(lines) < 2 {
		return false
	}

	text := strings.Join(lines, separator.Unix)
	openBrace := strings.Index(text, "{")
	closeBrace := strings.LastIndex(text, "}")

	if openBrace == -1 ||
		closeBrace == -1 ||
		closeBrace <= openBrace {
		return false
	}

	return strings.TrimSpace(text[openBrace+1:closeBrace]) == ""
}
