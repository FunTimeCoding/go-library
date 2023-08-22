package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func DeleteFile(name string) {
	errors.PanicOnError(os.Remove(name))
}
