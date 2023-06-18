package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func Home() string {
	result, e := os.UserHomeDir()
	errors.PanicOnError(e)

	return result
}
