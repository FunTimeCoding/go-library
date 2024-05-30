package system

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
)

func GoFiles(root string) []string {
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

				if !i.IsDir() && filepath.Ext(path) == constant.GoExtension {
					result = append(result, path)
				}

				return nil
			},
		),
	)

	return result
}
