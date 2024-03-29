package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/user"
)

func User() *user.User {
	result, e := user.Current()
	errors.PanicOnError(e)

	return result
}
