package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"path/filepath"
)

func RelativePath(
	base string,
	target string,
) string {
	result, e := filepath.Rel(base, target)
	errors.PanicOnError(e)

	return result
}
