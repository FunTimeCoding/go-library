package system

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"os"
	"path/filepath"
)

func CreateTarZip(
	sourceDirectory string,
	destinationFile string,
) {
	f := Create(destinationFile)
	defer errors.LogClose(f)

	zip := gzip.NewWriter(f)
	defer errors.LogClose(zip)

	writer := tar.NewWriter(zip)
	defer errors.LogClose(writer)

	errors.PanicOnError(
		filepath.Walk(
			sourceDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				errors.PanicOnError(e)

				h, headerFail := tar.FileInfoHeader(i, i.Name())
				errors.PanicOnError(headerFail)

				h.Name = filepath.ToSlash(
					join.Absolute(sourceDirectory, path),
				)
				errors.PanicOnError(writer.WriteHeader(h))

				if !i.Mode().IsRegular() {
					return nil
				}

				l := Open(path)
				Copy(l, writer)
				errors.PanicClose(l)

				return nil
			},
		),
	)
}
