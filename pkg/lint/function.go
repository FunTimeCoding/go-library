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
	result := file_report.New(path, r)
	var funcStart int
	var funcBraceDepth int
	var funcLines []string
	var inFuncBody bool

	for result.Scan() {
		line, number := result.Text()
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
				if isEmptyFunctionBody(funcLines) {
					result.AddConcern(
						constant.EmptyFunctionBodyKey,
						constant.EmptyFunctionBodyText,
						path,
						funcStart,
						strings.Join(funcLines, "\n"),
					)
				}
				inFuncBody = false
				funcLines = nil
			}
		}

		result.ChangedLine(line)
	}

	result.Fix = func() {
		if result.Fixed != "" {
			fmt.Printf("Fix functions %s\n", path)
			system.SaveFile(path, result.Fixed)
		}
	}

	return result.Finalize()
}

func isEmptyFunctionBody(lines []string) bool {
	if len(lines) < 2 {
		return false
	}

	fullText := strings.Join(lines, "\n")
	openBraceIdx := strings.Index(fullText, "{")
	closeBraceIdx := strings.LastIndex(fullText, "}")

	if openBraceIdx == -1 ||
		closeBraceIdx == -1 ||
		closeBraceIdx <= openBraceIdx {
		return false
	}

	body := fullText[openBraceIdx+1 : closeBraceIdx]

	return strings.TrimSpace(body) == ""
}
