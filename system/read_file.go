package file

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
	"path"
)

func ReadFile(name string) string {
	directory, directoryFail := os.Getwd()
	errors.FatalOnError(directoryFail)

	file, readFail := os.ReadFile(path.Join(directory, name))
	errors.FatalOnError(readFail)

	return string(file)
}
