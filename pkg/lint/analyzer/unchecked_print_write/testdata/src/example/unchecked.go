package example

import (
	"fmt"
	"os"
)

func Unchecked() {
	fmt.Fprintf(os.Stderr, "error: %s\n", "something") // want `use writer\.Print, errors\.Printf, or check the error from fmt\.Fprintf`
}
