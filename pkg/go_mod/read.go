package go_mod

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"golang.org/x/mod/modfile"
)

func Read() *modfile.File {
	result, e := modfile.Parse(ModFile, system.ReadBytes(ModFile), nil)
	errors.PanicOnError(e)

	return result
}
