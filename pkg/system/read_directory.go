package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ReadDirectory(name string) []os.DirEntry {
	result, e := os.ReadDir(name)
	errors.PanicOnError(e)

	return result
}
