package object_notation

import "github.com/funtimecoding/go-library/errors"

func DecodeStrict(
	value string,
	structure any,
) {
	errors.PanicOnError(Decode(value, structure))
}
