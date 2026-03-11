package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"strings"
	"unicode"
)

func TypeCount(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var count int
	var firstLine int
	var firstText string

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if !strings.HasPrefix(trimmed, "type ") {
			continue
		}

		rest := strings.TrimPrefix(trimmed, "type ")

		if rest == "" || !unicode.IsUpper(rune(rest[0])) {
			continue
		}

		count++

		if count == 1 {
			firstLine = number
			firstText = line
		}
	}

	if count > 1 {
		s.AddConcern(
			constant.MultipleExportedTypesKey,
			constant.MultipleExportedTypesText,
			path,
			firstLine,
			firstText,
		)
	}

	return s.Finalize()
}
