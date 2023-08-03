package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func ReadFile(name string) string {
	result, e := os.ReadFile(name)
	errors.PanicOnError(e)

	return string(result)
}
