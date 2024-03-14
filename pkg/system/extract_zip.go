package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
)

func ExtractZip(
	sourceFile string,
	destinationDirectory string,
) {
	f := Open(sourceFile)
	defer func() {
		errors.LogClose(f)
	}()

	for _, file := range ZipReader(sourceFile).File {
		ExtractZipFile(file, destinationDirectory)
	}
}
