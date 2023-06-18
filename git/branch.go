package git

import (
	"github.com/funtimecoding/go-library/errors"
	"strings"
)

func Branch(path string) string {
	repository := Open(path)
	branch, e := repository.Head()
	errors.PanicOnError(e)

	return strings.TrimPrefix(branch.Name().String(), "refs/heads/")
}
