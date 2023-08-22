package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func RemoveDirectory(path string) {
	errors.PanicOnError(os.RemoveAll(path))
}
