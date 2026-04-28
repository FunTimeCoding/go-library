package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func RemoveFile(path string) {
	errors.PanicOnError(os.Remove(path))
}
