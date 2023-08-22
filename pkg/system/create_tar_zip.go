package system

import (
	"archive/tar"
	"compress/gzip"
	errors2 "github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
)

func CreateTarZip(
	sourceDirectory string,
	destinationFile string,
) {
	f := Create(destinationFile)
	defer func() {
		errors2.LogClose(f)
	}()

	zip := gzip.NewWriter(f)
	defer func() {
		errors2.LogClose(zip)
	}()

	writer := tar.NewWriter(zip)
	defer func() {
		errors2.LogClose(writer)
	}()

	errors2.PanicOnError(
		filepath.Walk(
			sourceDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				errors2.PanicOnError(e)

				header, headerFail := tar.FileInfoHeader(i, i.Name())
				errors2.PanicOnError(headerFail)

				header.Name = filepath.ToSlash(Join(sourceDirectory, path))
				errors2.PanicOnError(writer.WriteHeader(header))

				if !i.Mode().IsRegular() {
					return nil
				}

				file := Open(path)
				defer func() {
					errors2.LogClose(file)
				}()

				Copy(file, writer)

				return nil
			},
		),
	)
}
