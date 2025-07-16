package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func MoveCopy(
	source string,
	destination string,
) {
	input := Open(source)
	output := Create(destination)
	defer errors.LogClose(output)
	Copy(input, output)
	// Close before trying to remove for Windows
	// https://stackoverflow.com/a/64943554/246801
	errors.PanicClose(input)
	errors.PanicOnError(os.Remove(source))
}
