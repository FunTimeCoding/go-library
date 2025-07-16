package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ChangeMode(
	path string,
	mode os.FileMode,
) {
	errors.PanicOnError(os.Chmod(path, mode))
}
