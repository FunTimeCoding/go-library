package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os"
	"path"
)

func ReadFile(name string) string {
	file, e := os.ReadFile(path.Join(WorkingDirectory(), name))
	errors.PanicOnError(e)

	return string(file)
}
