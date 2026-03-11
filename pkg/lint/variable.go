package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"regexp"
	"strings"
)

var errPattern = regexp.MustCompile(`\berr(?:\s*(?::=|=[^=])|,)`)

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

		if errPattern.MatchString(trimmed) {
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
