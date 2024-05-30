package lint

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/separator"
	"io"
	"strings"
)

func Fix(r io.Reader) string {
	var importBlock []string
	var insideImportBlock bool
	var modified bool
	s := bufio.NewScanner(r)
	var result strings.Builder

	for s.Scan() {
		line := s.Text()

		if !insideImportBlock && strings.HasPrefix(line, "import (") {
			insideImportBlock = true
		}

		if insideImportBlock {
			importBlock = append(importBlock, line)

			if strings.HasPrefix(line, ")") {
				insideImportBlock = false

				if len(importBlock) == 3 {
					result.WriteString(
						fmt.Sprintf(
							"import %s\n",
							strings.TrimSpace(importBlock[1]),
						),
					)
					modified = true
				} else {
					result.WriteString(
						strings.Join(
							importBlock,
							separator.Unix,
						) + separator.Unix,
					)
				}

				importBlock = importBlock[:0]
			}
		} else {
			result.WriteString(line + "\n")
		}
	}

	if modified {
		return result.String()
	}

	return ""
}
