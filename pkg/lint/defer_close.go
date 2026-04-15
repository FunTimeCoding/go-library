package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"regexp"
	"strings"
)

var deferClosePattern = regexp.MustCompile(
	`^(\s*)defer (\w+)\.Close\(\)$`,
)

func DeferClose(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var hasErrorsImport bool
	var replaced bool

	for s.Scan() {
		line, number := s.Text()

		if strings.Contains(line, constant.ErrorsImportPath) {
			hasErrorsImport = true
		}

		m := deferClosePattern.FindStringSubmatch(line)

		if m == nil {
			s.PassLine(line)

			continue
		}

		s.ChangedLine(
			fmt.Sprintf(
				"%sdefer errors.PanicClose(%s)",
				m[1],
				m[2],
			),
		)
		s.AddConcern(
			constant.DeferCloseKey,
			constant.DeferCloseText,
			path,
			number,
			line,
			true,
		)
		replaced = true
	}

	report := s.Finalize()

	if replaced && !hasErrorsImport && report.Fixed != "" {
		report.Fixed = addErrorsImport(report.Fixed)
	}

	return report
}
