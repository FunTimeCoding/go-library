package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func RemoveDirectory(path string) {
	errors.PanicOnError(os.RemoveAll(path))
}
