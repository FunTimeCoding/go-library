package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func fixtureFile(name string) (*os.File, error) {
	return os.Open(system.Join("..", "fixture", name))
}
