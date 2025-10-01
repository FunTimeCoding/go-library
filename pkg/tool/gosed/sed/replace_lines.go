package sed

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"strings"
)

func replaceLines(
	content string,
	replaces map[string]string,
) string {
	lines := split.NewLine(content)

	for i, l := range lines {
		for prefix, replace := range replaces {
			if strings.HasPrefix(strings.TrimSpace(l), prefix) {
				// Keep indentation
				indent := l[:len(l)-len(strings.TrimLeft(l, " \t"))]
				lines[i] = indent + prefix + replace
				break
			}
		}
	}

	return join.NewLine(lines)
}
