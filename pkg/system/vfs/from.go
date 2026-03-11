package vfs

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func From(dir string) *VFS {
	result := New()
	errors.PanicOnError(
		filepath.Walk(
			dir,
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
				relative := strings.TrimPrefix(path, dir+string(filepath.Separator))
				result.Write(relative, string(content))

				return nil
			},
		),
	)

	return result
}
