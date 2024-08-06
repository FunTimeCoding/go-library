package environment

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func Required(
	name string,
	exitCode int,
) string {
	if s := os.Getenv(name); s != "" {
		return s
	}

	system.Exitf(exitCode, "environment variable empty: %s\n", name)

	return ""
}
