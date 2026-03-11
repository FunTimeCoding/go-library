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
	var inBacktick bool
	var blockStack []bool // true = control block
	var pendingControl bool

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)
		isBlank := trimmed == ""

		wasInBacktick := inBacktick

		if strings.Count(line, "`")%2 == 1 {
			inBacktick = !inBacktick
		}

		if wasInBacktick {
			s.ChangedLine(line)
			prevLine = line
			prevWasBlank = isBlank

			continue
		}

		isControlStart := strings.HasPrefix(trimmed, "if ") ||
			strings.HasPrefix(trimmed, "for ") ||
			strings.HasPrefix(trimmed, "switch ") ||
			strings.HasPrefix(trimmed, "select ")

		isExit := trimmed == "return" || strings.HasPrefix(trimmed, "return ") ||
			trimmed == "break" || strings.HasPrefix(trimmed, "break ") ||
			trimmed == "continue" || strings.HasPrefix(trimmed, "continue ")

		isClosingBrace := strings.HasPrefix(trimmed, "}") ||
			strings.HasPrefix(trimmed, "case ") ||
			trimmed == "default:"

		prevTrimmed := strings.TrimSpace(prevLine)
		prevOpensBlock := strings.HasSuffix(prevTrimmed, "{") ||
			strings.HasPrefix(prevTrimmed, "case ") ||
			prevTrimmed == "default:" ||
			strings.HasPrefix(prevTrimmed, "//")

		isElseContinuation := strings.HasPrefix(trimmed, "} else")
		endsWithBrace := strings.HasSuffix(trimmed, "{")

		if isElseContinuation {
			if len(blockStack) > 0 {
				blockStack = blockStack[:len(blockStack)-1]
			}

			if endsWithBrace {
				blockStack = append(blockStack, true)
			}

			pendingControl = false
		} else if strings.HasPrefix(trimmed, "}") && endsWithBrace {
			if len(blockStack) > 0 {
				blockStack = blockStack[:len(blockStack)-1]
			}

			blockStack = append(blockStack, false)
			pendingControl = false
		} else if endsWithBrace {
			blockStack = append(blockStack, isControlStart || pendingControl)
			pendingControl = false
		} else if isControlStart {
			pendingControl = true
		} else if !isBlank && !isClosingBrace {
			pendingControl = false
		}

		if isControlStart && !prevWasBlank && prevLine != "" && !prevOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBeforeControlKey,
				constant.MissingBlankBeforeControlText,
				path,
				number,
				line,
			)
			needBlankAfterClosingBrace = false
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
			needBlankAfterClosingBrace = false
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

		if trimmed == "}" {
			isControl := false

			if len(blockStack) > 0 {
				isControl = blockStack[len(blockStack)-1]
				blockStack = blockStack[:len(blockStack)-1]
			}

			needBlankAfterClosingBrace = isControl
		} else {
			needBlankAfterClosingBrace = false
		}

		prevLine = line
		prevWasBlank = isBlank
	}

	return s.Finalize()
}
