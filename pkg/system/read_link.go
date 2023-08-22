package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ReadLink(name string) string {
	result, e := os.Readlink(name)
	errors.PanicOnError(e)

	return result
}
