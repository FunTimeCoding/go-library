package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Stat(name string) os.FileInfo {
	result, e := os.Stat(name)
	errors.PanicOnError(e)

	return result
}
