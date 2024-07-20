package environment

import (
	"fmt"
	"os"
)

func Get(
	name string,
	exitCode int,
) string {
	result := os.Getenv(name)

	if result == "" {
		fmt.Printf("%s not set\n", name)

		os.Exit(exitCode)
	}

	return result
}
