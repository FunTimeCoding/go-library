package system

import "github.com/funtimecoding/go-library/pkg/errors"

func CleanAddressStrict(s string) string {
	clean, e := CleanAddress(s)
	errors.PanicOnError(e)

	return clean
}
