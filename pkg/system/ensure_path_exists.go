package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func EnsurePathExists(path string) {
	if _, e := os.Stat(path); os.IsNotExist(e) {
		errors.LogOnError(os.Mkdir(path, 0755))
	}
}
