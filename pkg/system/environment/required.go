package environment

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func Required(
	name string,
	exit int,
) string {
	if s := os.Getenv(name); s != "" {
		return s
	}

	system.Exitf(exit, "environment variable empty: %s\n", name)

	return ""
}
