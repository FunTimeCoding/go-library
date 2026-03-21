package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"strings"
)

func FunctionCount(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)

	if strings.HasSuffix(path, "_test.go") {
		for s.Scan() {
			line, _ := s.Text()
			s.PassLine(line)
		}

		return s.Finalize()
	}

	var count int
	var firstLine int
	var firstText string
	var depth int

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "func ") && depth == 0 {
			count++

			if count == 1 {
				firstLine = number
				firstText = line
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
			constant.MultipleFunctionsKey,
			constant.MultipleFunctionsText,
			path,
			firstLine,
			firstText,
		)
	}

	return s.Finalize()
}
