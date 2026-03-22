package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"strings"
)

func ForbiddenImport(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var insideBlock bool

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(line, "import (") {
			insideBlock = true
			s.PassLine(line)

			continue
		}

		if insideBlock {
			if strings.HasPrefix(line, ")") {
				insideBlock = false
			} else if trimmed == `"flag"` {
				s.AddConcern(
					constant.ForbiddenImportKey,
					constant.ForbiddenImportText,
					path,
					number,
					line,
					false,
				)
			}

			s.PassLine(line)

			continue
		}

		if line == `import "flag"` {
			s.AddConcern(
				constant.ForbiddenImportKey,
				constant.ForbiddenImportText,
				path,
				number,
				line,
				false,
			)
		}

		s.PassLine(line)
	}

	return s.Finalize()
}
