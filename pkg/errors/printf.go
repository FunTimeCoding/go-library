package errors

import (
	"fmt"
	"os"
)

func Printf(
	format string,
	a ...any,
) int {
	result, e := fmt.Fprintf(os.Stderr, format, a...)
	PanicOnError(e)

	return result
}
