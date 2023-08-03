package system

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/go-library/errors"
	"os"
	"path/filepath"
)

func CreateTarZip(
	sourceDirectory string,
	destinationFile string,
) {
	f := Create(destinationFile)
	defer func() {
		errors.LogClose(f)
	}()

	zip := gzip.NewWriter(f)
	defer func() {
		errors.LogClose(zip)
	}()

	writer := tar.NewWriter(zip)
	defer func() {
		errors.LogClose(writer)
	}()

	errors.PanicOnError(
		filepath.Walk(
			sourceDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				errors.PanicOnError(e)

				header, headerFail := tar.FileInfoHeader(i, i.Name())
				errors.PanicOnError(headerFail)

				header.Name = filepath.ToSlash(
					filepath.Join(sourceDirectory, path),
				)
				errors.PanicOnError(writer.WriteHeader(header))

				if !i.Mode().IsRegular() {
					return nil
				}

				file := Open(path)
				defer func() {
					errors.LogClose(file)
				}()

				Copy(file, writer)

				return nil
			},
		),
	)
}
