package environment

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Set(
	k string,
	v string,
) {
	errors.PanicOnError(os.Setenv(k, v))
}
