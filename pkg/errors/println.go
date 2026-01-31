package errors

import (
	"fmt"
	"os"
)

func Println(a ...any) int {
	result, e := fmt.Fprintln(os.Stderr, a...)
	PanicOnError(e)

	return result
}
