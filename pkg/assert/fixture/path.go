package fixture

import (
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func Path(path ...string) string {
	return project.FixturePath(join.Relative(path...))
}
