package status

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/repository"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
	"path/filepath"
	"strings"
)

func collect(maxDepth int) []*repository.Repository {
	root := system.WorkingDirectory()
	var result []*repository.Repository
	errors.PanicOnError(
		filepath.Walk(
			root,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return nil
				}

				relative, _ := filepath.Rel(root, path)
				depth := strings.Count(relative, string(os.PathSeparator))

				if depth > maxDepth {
					if i.IsDir() {
						return filepath.SkipDir
					}

					return nil
				}

				if i.IsDir() && i.Name() == constant.Directory {
					result = append(
						result,
						checkRepository(filepath.Dir(path)),
					)

					return filepath.SkipDir
				}

				return nil
			},
		),
	)

	return result
}
