package markup

import "github.com/funtimecoding/go-library/pkg/errors"

func MustDecode(
	value string,
	structure any,
) {
	errors.PanicOnError(Decode(value, structure))
}
