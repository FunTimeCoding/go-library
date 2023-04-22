package hypertext

import (
	"os"
	"path/filepath"
)

func fixtureFile(name string) (*os.File, error) {
	return os.Open(filepath.Join("..", "fixture", name))
}
