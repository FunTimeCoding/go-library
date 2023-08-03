package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func Open(path string) *os.File {
	result, e := os.Open(path)
	errors.PanicOnError(e)

	return result
}
