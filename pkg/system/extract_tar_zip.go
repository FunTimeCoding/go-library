package system

import (
	"archive/tar"
	errors2 "github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"os"
)

func ExtractTarZip(
	sourceFile string,
	destinationDirectory string,
) {
	f := Open(sourceFile)
	defer func() {
		errors2.LogClose(f)
	}()

	zip := ZipReader(f)
	defer func() {
		errors2.LogClose(zip)
	}()

	reader := tar.NewReader(zip)

	for {
		header, nextFail := reader.Next()

		if nextFail == io.EOF {
			break
		}

		errors2.PanicOnError(nextFail)
		target := Join(destinationDirectory, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			MakeDirectory(target)
		case tar.TypeReg:
			file := OpenFile(target, os.FileMode(header.Mode))

			if _, g := io.Copy(file, reader); g != nil {
				errors2.PanicClose(file)
				errors2.PanicOnError(g)
			}

			errors2.PanicClose(file)
		}
	}
}
