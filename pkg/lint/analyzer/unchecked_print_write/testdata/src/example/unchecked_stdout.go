package example

import (
	"fmt"
	"os"
)

func UncheckedStdout() {
	fmt.Fprintf(os.Stdout, "output: %s\n", "something") // want `use writer\.Print, errors\.Printf, or check the error from fmt\.Fprintf`
}
