package system

import "github.com/funtimecoding/go-library/pkg/errors"

func Run(s ...string) string {
	result, e := RunError(s...)
	errors.PanicOnError(e)

	return result
}
