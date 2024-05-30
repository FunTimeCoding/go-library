package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"os"
	"path/filepath"
	"strings"
)

func EmptyDirectories(root string) []string {
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

				if i.IsDir() {
					if !strings.Contains(path, constant.Directory) {
						if IsEmptyDirectory(path) {
							result = append(result, path)
						}
					}
				}

				return nil
			},
		),
	)

	return result
}
