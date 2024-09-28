package text

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/text/optimize_settings"
	"strings"
)

func OptimizeWhitespace(
	input string,
	s *optimize_settings.Settings,
) string {
	if s == nil {
		s = optimize_settings.New()
	}

	var result []string
	var blankLines int

	for _, line := range split.NewLine(input) {
		line = strings.TrimSpace(line)

		if line == "" {
			blankLines++

			if blankLines > s.AllowedBlankLines {
				continue
			}
		} else {
			blankLines = 0
		}

		result = append(result, line)
	}

	if s.NewlineAtEnd {
		// If no newline at the end, add one
		if len(result) > 0 && result[len(result)-1] != "" {
			result = append(result, "")
		}
	}

	return join.NewLine(result)
}
