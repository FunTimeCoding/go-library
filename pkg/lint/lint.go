package lint

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"strings"
)

func Lint(r io.Reader) string {
	var importBlock []string
	var insideImportBlock bool
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()

		if strings.HasPrefix(line, "import (") {
			insideImportBlock = true
			importBlock = append(importBlock, line)

			continue
		}

		if insideImportBlock {
			if strings.HasPrefix(line, ")") {
				importBlock = append(importBlock, line)
				insideImportBlock = false

				if len(importBlock) == 3 {
					return fmt.Sprintf(
						"import %s",
						strings.TrimSpace(importBlock[1]),
					)
				}

				importBlock = nil
			} else {
				importBlock = append(importBlock, line)
			}

			continue
		}
	}

	errors.PanicOnError(s.Err())

	return ""
}
