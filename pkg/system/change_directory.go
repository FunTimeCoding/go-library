package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ChangeDirectory(path string) {
	errors.PanicOnError(os.Chdir(path))
}
