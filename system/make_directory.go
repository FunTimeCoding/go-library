package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func MakeDirectory(path string) {
	errors.PanicOnError(os.MkdirAll(path, 0755))
}
