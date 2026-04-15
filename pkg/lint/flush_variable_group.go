package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"strings"
)

func flushVariableGroup(
	s *file_report.Report,
	lines []string,
	path string,
	startLine int,
) {
	s.ChangedLine("var (")

	for _, line := range lines {
		s.ChangedLine("\t" + strings.TrimPrefix(line, "var "))
	}

	s.ChangedLine(")")
	s.AddConcern(
		constant.VariableGroupingKey,
		constant.VariableGroupingText,
		path,
		startLine,
		lines[0],
		true,
	)
}
