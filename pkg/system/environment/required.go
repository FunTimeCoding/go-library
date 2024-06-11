package environment

import (
	"fmt"
	"os"
)

func Required(
	name string,
	exitCode int,
) string {
	if s := os.Getenv(name); s != "" {
		return s
	}

	fmt.Printf("environment variable empty: %s\n", name)
	os.Exit(exitCode)

	return ""
}
