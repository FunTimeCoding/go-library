package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Executable(path string) {
	errors.PanicOnError(os.Chmod(path, Stat(path).Mode()|0111))
}
