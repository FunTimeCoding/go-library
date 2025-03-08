package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
)

func FilesRecursive(root string) []string {
	var result []string
	errors.PanicOnError(
		filepath.Walk(
			root,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if !i.IsDir() {
					result = append(result, path)
				}

				return nil
			},
		),
	)

	return result
}
