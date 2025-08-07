package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"io"
	"slices"
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
	var reportedBlank bool

	for result.Scan() {
		line, number := result.Text()

		if !inside && strings.HasPrefix(line, "import (") {
			inside = true
			start = number
			blockText = line
		}

		if inside {
			if line == "" {
				if !reportedBlank {
					reportedBlank = true
					result.AddConcern(
						constant.ImportBlankKey,
						constant.ImportBlankText,
						path,
						start,
						blockText,
					)
				}

				continue
			}

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
						constant.SingleMultiImportKey,
						constant.SingleMultiImportText,
						path,
						start,
						blockText,
					)
				} else {
					sorted := make([]string, len(block)-2)
					copy(sorted, block[1:len(block)-1])
					SortImports(sorted)
					sorted = append([]string{"import ("}, sorted...)
					sorted = append(sorted, ")")

					if !slices.Equal(block, sorted) {
						result.AddConcern(
							constant.UnsortedImportsKey,
							constant.UnsortedImportsText,
							path,
							start,
							blockText,
						)
					}

					result.ChangedLine(join.NewLine(sorted))
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
