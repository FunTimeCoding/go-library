package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
	"strings"
)

func Spacing(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var prevLine string
	var prevWasBlank bool
	var needBlankAfterClosingBrace bool

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)
		isBlank := trimmed == ""

		isControlStart := strings.HasPrefix(trimmed, "if ") ||
			strings.HasPrefix(trimmed, "for ") ||
			strings.HasPrefix(trimmed, "switch ") ||
			strings.HasPrefix(trimmed, "select ")

		isExit := strings.HasPrefix(trimmed, "return") ||
			strings.HasPrefix(trimmed, "break") ||
			strings.HasPrefix(trimmed, "continue")

		isClosingBrace := trimmed == "}" ||
			strings.HasPrefix(trimmed, "} else")

		prevOpensBlock := strings.HasSuffix(
			strings.TrimSpace(prevLine), "{",
		)

		if isControlStart && !prevWasBlank && prevLine != "" && !prevOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBeforeControlKey,
				constant.MissingBlankBeforeControlText,
				path,
				number,
				line,
			)
		}

		if isExit && !prevWasBlank && prevLine != "" && !prevOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBeforeExitKey,
				constant.MissingBlankBeforeExitText,
				path,
				number,
				line,
			)
		}

		if needBlankAfterClosingBrace && !isBlank && !isClosingBrace {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankAfterControlKey,
				constant.MissingBlankAfterControlText,
				path,
				number,
				line,
			)
			needBlankAfterClosingBrace = false
		}

		if isBlank && prevWasBlank {
			s.AddConcern(
				constant.ExtraneousBlankLineKey,
				constant.ExtraneousBlankLineText,
				path,
				number,
				line,
			)
			needBlankAfterClosingBrace = false

			continue
		}

		s.ChangedLine(line)

		if isClosingBrace && !strings.HasSuffix(trimmed, "{") {
			needBlankAfterClosingBrace = true
		} else if !isBlank {
			needBlankAfterClosingBrace = false
		}

		prevLine = line
		prevWasBlank = isBlank
	}

	return s.Finalize()
}
