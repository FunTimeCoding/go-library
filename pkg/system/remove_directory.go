package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Remove(path string) {
	errors.PanicOnError(os.RemoveAll(path))
}
