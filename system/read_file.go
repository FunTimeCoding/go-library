package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
)

func ReadFile(name string) string {
	file, e := os.ReadFile(Join(WorkingDirectory(), name))
	errors.PanicOnError(e)

	return string(file)
}
