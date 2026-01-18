package system

import (
	"archive/tar"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func TarWriteHeader(
	w *tar.Writer,
	h *tar.Header,
) {
	errors.PanicOnError(w.WriteHeader(h))
}
