package system

import (
	"github.com/funtimecoding/go-library/pkg/system/join"
	"os"
)

func OpenHome(subPath string) *os.File {
	return Open(join.Absolute(Home(), subPath))
}
