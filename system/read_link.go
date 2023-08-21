package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func ReadLink(name string) string {
	result, e := os.Readlink(name)
	errors.PanicOnError(e)

	return result
}
