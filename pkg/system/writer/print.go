package writer

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func Print(
	w io.Writer,
	format string,
	a ...any,
) {
	_, err := fmt.Fprintf(w, format, a...)
	errors.PanicOnError(err)
}
