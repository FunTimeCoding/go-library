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
	result := file_report.New(path, r)

	for result.Scan() {
		line, number := result.Text()

		if number == 1 {
			if line != constant.FrontMatterDelimiter {
				result.ChangedLine(constant.FrontMatterDelimiter)
				result.ChangedLine(line)
				result.AddConcern(
					constant.FrontMatterDelimiterKey,
					constant.FrontMatterDelimiterText,
					number,
					line,
				)
			}
		} else {
			result.ChangedLine(line)
		}
	}

	result.Fix = func() {
		if result.Fixed != "" {
			fmt.Printf(
				"Add front matter delimiter %s %+v\n",
				path,
				result,
			)
			system.SaveFile(path, result.Fixed)
		}
	}

	return result.Finalize()
}
