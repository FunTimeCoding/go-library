package system

import (
	"archive/zip"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ZipAdd(
	w *zip.Writer,
	f *os.File,
	path string,
) {
	write, e := w.Create(path)
	errors.PanicOnError(e)
	Copy(f, write)
}
