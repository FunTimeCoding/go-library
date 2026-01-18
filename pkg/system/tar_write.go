package system

import (
	"archive/tar"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func TarWrite(
	t *tar.Writer,
	b []byte,
) int {
	result, e := t.Write(b)
	errors.PanicOnError(e)

	return result
}

