package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"regexp"
	"strings"
)

var errPattern = regexp.MustCompile(`\berr(?:\s*(?::=|=[^=])|,)`)
var stringLiteral = regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)

func Variable(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "//") {
			continue
		}

		stripped := stringLiteral.ReplaceAllString(trimmed, `""`)

		if errPattern.MatchString(stripped) {
			s.AddConcern(
				constant.ErrVariableKey,
				constant.ErrVariableText,
				path,
				number,
				line,
			)
		}
	}

	return s.Finalize()
}
