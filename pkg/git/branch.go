package git

import "github.com/funtimecoding/go-library/pkg/errors"

func Branch(path string) string {
	b, e := Open(path).Head()
	errors.PanicOnError(e)

	return b.Name().Short()
}
