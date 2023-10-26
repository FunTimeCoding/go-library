package main

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	errors.PanicOnError(
		filepath.Walk(
			".",
			func(
				path string,
				info os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if info.IsDir() || filepath.Ext(path) != ".go" {
					return nil
				}

				file, openFail := os.Open(path)

				if openFail != nil {
					return openFail
				}

				defer func() {
					errors.LogOnError(file.Close())
				}()

				scanner := bufio.NewScanner(file)
				var importBlock []string
				var insideImportBlock bool

				for scanner.Scan() {
					line := scanner.Text()

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
								fmt.Printf(
									"Simplify import in %s: %s\n",
									path,
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

				if f := scanner.Err(); f != nil {
					return f
				}

				return nil
			},
		),
	)
}
