package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/lint"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
	"path/filepath"
)

func main() {
	errors.PanicOnError(
		filepath.Walk(
			".",
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if i.IsDir() || filepath.Ext(path) != ".go" {
					return nil
				}

				f := system.Open(path)
				defer errors.LogClose(f)

				if false {
					if linted := lint.Lint(f); linted != "" {
						fmt.Printf(
							"Simplify import in %s: %s\n",
							path,
							linted,
						)
					}
				}

				if linted := lint.Fix(f); linted != "" {
					fmt.Printf(
						"Simplify import in %s: %s\n",
						path,
						linted,
					)
					system.SaveFile(path, linted)
				}

				return nil
			},
		),
	)
}
