package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func Copy(
	source io.Reader,
	destination io.Writer,
) {
	_, e := io.Copy(destination, source)
	errors.PanicOnError(e)
}
