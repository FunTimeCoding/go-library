package environment

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func EnsureUnset(name string) {
	if os.Getenv(name) != "" {
		errors.PanicOnError(os.Unsetenv(name))
	}
}
