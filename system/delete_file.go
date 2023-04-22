package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func DeleteFile(name string) {
	errors.PanicOnError(os.Remove(name))
}
