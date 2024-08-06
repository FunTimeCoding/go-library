package environment

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func Get(
	name string,
	exitCode int,
) string {
	result := os.Getenv(name)

	if result == "" {
		system.Exitf(exitCode, "%s not set\n", name)
	}

	return result
}
