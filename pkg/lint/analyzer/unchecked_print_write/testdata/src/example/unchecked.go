package example

import (
	"fmt"
	"os"
)

func Unchecked() {
	fmt.Fprintf( // want `use writer\.Print, errors\.Printf, or check the error from fmt\.Fprintf`
		os.Stderr,
		"error: %s\n",
		"something",
	)
}
