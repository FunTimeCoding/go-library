package text

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"strings"
)

func OptimizeWhitespace(input string) string {
	var result []string
	var blankLines int

	for _, line := range split.NewLine(input) {
		line = strings.TrimSpace(line)

		if line == "" {
			blankLines++

			if blankLines > 1 {
				continue
			}
		} else {
			blankLines = 0
		}

		result = append(result, line)
	}

	return join.NewLine(result)
}
