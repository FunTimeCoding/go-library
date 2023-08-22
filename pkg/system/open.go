package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Open(path string) *os.File {
	result, e := os.Open(path)
	errors.PanicOnError(e)

	return result
}
