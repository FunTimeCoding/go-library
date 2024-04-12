package mime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func Write(
	part io.Writer,
	f *File,
) {
	_, e := part.Write(f.Data)
	errors.PanicOnError(e)
}
