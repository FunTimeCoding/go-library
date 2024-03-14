package system

import (
	"archive/zip"
	"github.com/funtimecoding/go-library/pkg/errors"
	"path/filepath"
)

func ExtractZipFile(
	f *zip.File,
	directory string,
) {
	reader, e := f.Open()
	errors.PanicOnError(e)

	defer func() {
		errors.PanicClose(reader)
	}()

	out := Create(filepath.Join(directory, f.Name))

	defer func() {
		errors.PanicClose(out)
	}()

	Copy(reader, out)
}
