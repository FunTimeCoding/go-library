package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
	"path"
)

func ReadFile(name string) string {
	directory, directoryFail := os.Getwd()
	errors.PanicOnError(directoryFail)

	file, readFail := os.ReadFile(path.Join(directory, name))
	errors.PanicOnError(readFail)

	return string(file)
}
