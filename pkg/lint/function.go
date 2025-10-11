package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/system"
	"io"
	"strings"
)

func Function(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var funcStart int
	var funcBraceDepth int
	var funcLines []string
	var inFuncBody bool

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "func ") && strings.Contains(
			line,
			"{",
		) {
			funcStart = number
			funcBraceDepth = strings.Count(line, "{") - strings.Count(
				line,
				"}",
			)
			funcLines = []string{line}
			inFuncBody = funcBraceDepth > 0
		} else if inFuncBody {
			funcLines = append(funcLines, line)
			funcBraceDepth += strings.Count(line, "{") - strings.Count(
				line,
				"}",
			)

			if funcBraceDepth == 0 {
				if emptyFunctionBody(funcLines) {
					fixedFunction := toSingleLine(funcLines)
					s.ChangedLine(fixedFunction)
					s.AddConcern(
						constant.EmptyFunctionBodyKey,
						constant.EmptyFunctionBodyText,
						path,
						funcStart,
						strings.Join(funcLines, "\n"),
					)
				} else {
					for _, funcLine := range funcLines {
						s.ChangedLine(funcLine)
					}
				}

				inFuncBody = false
				funcLines = nil

				continue
			}
		}

		if !inFuncBody {
			s.ChangedLine(line)
		}
	}

	s.Fix = func() {
		if s.Fixed != "" {
			fmt.Printf("Fix functions %s\n", path)
			system.SaveFile(path, s.Fixed)
		}
	}

	return s.Finalize()
}
