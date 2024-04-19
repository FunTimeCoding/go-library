package mime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func Write(
	w io.Writer,
	b []byte,
) {
	_, e := w.Write(b)
	errors.PanicOnError(e)
}
