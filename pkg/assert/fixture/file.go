package fixture

import (
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"os"
)

func File(path ...string) *os.File {
	return system.Open(project.FixturePath(join.Relative(path...)))
}
