package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Move(from string, to string) {
	errors.PanicOnError(os.Rename(from, to))
}
