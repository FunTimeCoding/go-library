package system

import (
	"github.com/funtimecoding/go-library/errors"
	"io"
)

func Copy(source io.Reader, destination io.Writer) {
	_, e := io.Copy(destination, source)
	errors.PanicOnError(e)
}
