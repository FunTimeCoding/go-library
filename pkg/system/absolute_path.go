package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"path/filepath"
)

func AbsolutePath(p string) string {
	result, e := filepath.Abs(p)
	errors.PanicOnError(e)

	return result
}
