package environment

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func ExitCode(
	name string,
	code int,
) string {
	result := os.Getenv(name)

	if result == "" {
		system.Exitf(code, "%s not set\n", name)
	}

	return result
}
