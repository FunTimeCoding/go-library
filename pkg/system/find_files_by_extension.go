package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
)

func FindFilesByExtensions(
	root string,
	extension string,
) []string {
	var result []string

	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				path string,
				d os.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				if !d.IsDir() && filepath.Ext(path) == extension {
					result = append(result, path)
				}

				return nil
			},
		),
	)

	return result
}
