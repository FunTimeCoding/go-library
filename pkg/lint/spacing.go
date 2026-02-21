package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/system"
	"io"
	"strings"
)

const (
	missingBlankBeforeControlKey  = "missing_blank_before_control"
	missingBlankBeforeControlText = "Missing blank line before control block"
	missingBlankBeforeExitKey     = "missing_blank_before_exit"
	missingBlankBeforeExitText    = "Missing blank line before return/break/continue"
	extraneousBlankLineKey        = "extraneous_blank_line"
	extraneousBlankLineText       = "Extraneous blank line"
	missingBlankAfterControlKey   = "missing_blank_after_control"
	missingBlankAfterControlText  = "Missing blank line after control block"
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
				missingBlankBeforeControlKey,
				missingBlankBeforeControlText,
				path,
				number,
				line,
			)
		}

		if isExit && !prevWasBlank && prevLine != "" && !prevOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				missingBlankBeforeExitKey,
				missingBlankBeforeExitText,
				path,
				number,
				line,
			)
		}

		if needBlankAfterClosingBrace && !isBlank && !isClosingBrace {
			s.ChangedLine("")
			s.AddConcern(
				missingBlankAfterControlKey,
				missingBlankAfterControlText,
				path,
				number,
				line,
			)
			needBlankAfterClosingBrace = false
		}

		if isBlank && prevWasBlank {
			s.AddConcern(
				extraneousBlankLineKey,
				extraneousBlankLineText,
				path,
				number,
				line,
			)

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

	s.Fix = func() {
		if s.Fixed != "" {
			fmt.Printf("Fix spacing %s\n", path)
			system.SaveFile(path, s.Fixed)
		}
	}

	return s.Finalize()
}
