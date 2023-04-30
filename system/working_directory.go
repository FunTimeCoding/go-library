package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func WorkingDirectory() string {
	d, e := os.Getwd()
	errors.PanicOnError(e)

	return d
}
