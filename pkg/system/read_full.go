package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func ReadFull(
	r io.Reader,
	b []byte,
) int {
	result, e := io.ReadFull(r, b)
	errors.PanicOnError(e)

	return result
}

