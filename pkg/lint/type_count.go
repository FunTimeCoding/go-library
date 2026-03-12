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
	var depth int

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "type ") {
			rest := strings.TrimPrefix(trimmed, "type ")

			if rest != "" && unicode.IsUpper(rune(rest[0])) && depth == 0 {
				count++

				if count == 1 {
					firstLine = number
					firstText = line
				}
			}
		}

		for _, c := range line {
			switch c {
			case '{':
				depth++
			case '}':
				depth--
			}
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
