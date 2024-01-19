package go_mod

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"golang.org/x/mod/modfile"
	"os"
)

func Read() *modfile.File {
	modBytes, readFail := os.ReadFile(ModFile)
	errors.PanicOnError(readFail)

	result, e := modfile.Parse(ModFile, modBytes, nil)
	errors.PanicOnError(e)

	return result
}
