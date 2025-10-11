package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"path/filepath"
)

func Glob(p string) []string {
	result, e := filepath.Glob(p)
	errors.PanicOnError(e)

	return result
}
