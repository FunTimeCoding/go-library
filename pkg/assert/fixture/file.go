package fixture

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func File(path ...string) *os.File {
	return system.Open(Path(path...))
}
