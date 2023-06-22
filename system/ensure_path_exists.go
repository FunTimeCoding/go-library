package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
	"path/filepath"
)

func EnsurePathExists(path string) {
	path = filepath.Dir(path)

	if _, e := os.Stat(path); os.IsNotExist(e) {
		errors.LogOnError(os.Mkdir(path, 0755))
	}
}
