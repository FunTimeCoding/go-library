package git

import (
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Branch(path string) string {
	branch, e := Open(path).Head()
	errors.PanicOnError(e)

	return branch.Name().Short()
}
