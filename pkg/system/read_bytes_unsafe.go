package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ReadBytesUnsafe(path string) []byte {
	result, e := os.ReadFile(path)
	errors.PanicOnError(e)

	return result
}
