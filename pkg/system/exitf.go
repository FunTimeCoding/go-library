package system

import (
	"fmt"
	"os"
)

func Exitf(
	code int,
	format string,
	a ...any,
) {
	fmt.Printf(format, a...)
	os.Exit(code)
}
