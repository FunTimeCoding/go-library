package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Root(name string) *os.Root {
	result, e := os.OpenRoot(name)
	errors.PanicOnError(e)

	return result
}
