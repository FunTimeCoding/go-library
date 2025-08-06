package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func ExecutablePath() string {
	result, e := os.Executable()
	errors.PanicOnError(e)

	return result
}
