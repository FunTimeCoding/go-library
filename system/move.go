package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func Move(from string, to string) {
	errors.PanicOnError(os.Rename(from, to))
}
