package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ReadBytes(name string) []byte {
	result, e := os.ReadFile(name)
	errors.PanicOnError(e)

	return result
}
