package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func Open(path string) *os.File {
	file, e := os.Open(path)
	errors.PanicOnError(e)

	return file
}
