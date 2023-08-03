package system

import (
	"archive/tar"
	"github.com/funtimecoding/go-library/errors"
	"io"
	"os"
	"path/filepath"
)

func ExtractTarZip(
	sourceFile string,
	destinationDirectory string,
) {
	f := Open(sourceFile)
	defer func() {
		errors.LogClose(f)
	}()

	zip := ZipReader(f)
	defer func() {
		errors.LogClose(zip)
	}()

	reader := tar.NewReader(zip)

	for {
		header, nextFail := reader.Next()

		if nextFail == io.EOF {
			break
		}

		errors.PanicOnError(nextFail)
		target := filepath.Join(destinationDirectory, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			MakeDirectory(target)
		case tar.TypeReg:
			file := OpenFile(target, os.FileMode(header.Mode))

			if _, g := io.Copy(file, reader); g != nil {
				errors.PanicClose(file)
				errors.PanicOnError(g)
			}

			errors.PanicClose(file)
		}
	}
}
