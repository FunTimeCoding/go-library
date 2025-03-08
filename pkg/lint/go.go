package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"io"
	"strings"
)

func Go(
	path string,
	r io.Reader,
) *file_report.Report {
	result := file_report.New(path, r)
	var block []string
	var inside bool
	var start int
	var blockText string

	for result.Scan() {
		line, number := result.Text()

		if !inside && strings.HasPrefix(line, "import (") {
			inside = true
			start = number
			blockText = line
		}

		if inside {
			block = append(block, line)

			if strings.HasPrefix(line, ")") {
				inside = false

				if len(block) == 3 {
					result.ChangedLine(
						fmt.Sprintf(
							"import %s",
							strings.TrimSpace(block[1]),
						),
					)
					result.AddConcern(
						constant.SingleMultiLineImportKey,
						constant.SingleMultiLineImportText,
						path,
						start,
						blockText,
					)
				} else {
					result.ChangedLine(join.NewLine(block))
				}

				block = block[:0]
			}
		} else {
			result.ChangedLine(line)
		}
	}

	result.Fix = func() {
		if result.Fixed != "" {
			fmt.Printf("Simplify import %s\n", path)
			system.SaveFile(path, result.Fixed)
		}
	}

	return result.Finalize()
}
