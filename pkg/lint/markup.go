package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/system"
	"io"
)

func Markup(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)

	for s.Scan() {
		line, number := s.Text()

		if number == 1 {
			if line != constant.FrontMatterDelimiter {
				s.ChangedLine(constant.FrontMatterDelimiter)
				s.ChangedLine(line)
				s.AddConcern(
					constant.FrontMatterDelimiterKey,
					constant.FrontMatterDelimiterText,
					path,
					number,
					line,
				)
			}
		} else {
			s.ChangedLine(line)
		}
	}

	s.Fix = func() {
		if s.Fixed != "" {
			fmt.Printf("Add front matter delimiter %s\n", path)
			system.SaveFile(path, s.Fixed)
		}
	}

	return s.Finalize()
}
