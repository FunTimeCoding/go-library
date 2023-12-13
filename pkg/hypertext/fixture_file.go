package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func fixtureFile(name string) *os.File {
	result, e := os.Open(system.Join("..", "fixture", name))
	errors.PanicOnError(e)

	return result
}
