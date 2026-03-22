package virtual_file_system

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func From(directory string) *System {
	result := New()
	errors.PanicOnError(
		filepath.Walk(
			directory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if i.IsDir() {
					return nil
				}

				content, e := os.ReadFile(path)
				errors.PanicOnError(e)
				relative := strings.TrimPrefix(
					path,
					fmt.Sprintf(
						"%s%s",
						directory,
						string(filepath.Separator),
					),
				)
				result.Write(relative, string(content))

				return nil
			},
		),
	)

	return result
}
