package system

import (
	"os"
	"path"
)

func OpenHome(subPath string) *os.File {
	return Open(path.Join(Home(), subPath))
}
