package example

import (
	"fmt"
	"os"
)

func UncheckedStdout() {
	fmt.Fprintf( // want `use writer\.Print, errors\.Printf, or check the error from fmt\.Fprintf`
		os.Stdout,
		"output: %s\n",
		"something",
	)
}
