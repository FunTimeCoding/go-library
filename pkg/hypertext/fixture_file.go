package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"os"
)

func fixtureFile(name string) *os.File {
	result, e := os.Open(system.Join("..", constant.FixturePath, name))
	errors.PanicOnError(e)

	return result
}
