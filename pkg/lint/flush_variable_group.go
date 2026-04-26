package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/strings/join"
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
		s.ChangedLine(join.Empty("\t", strings.TrimPrefix(line, "var ")))
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
