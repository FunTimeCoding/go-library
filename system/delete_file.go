package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func DeleteFile(name string) {
	errors.FatalOnError(os.Remove(name))
}
