package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func ReadAll(r io.Reader) []byte {
	result, e := io.ReadAll(r)
	errors.PanicOnError(e)

	return result
}
