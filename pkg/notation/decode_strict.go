package notation

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func DecodeStrict(
	value string,
	structure any,
	verbose bool,
) {
	e := Decode(value, structure)

	if verbose && e != nil {
		fmt.Printf("Could not decode: %s\n", value)
	}

	errors.PanicOnError(e)
}
