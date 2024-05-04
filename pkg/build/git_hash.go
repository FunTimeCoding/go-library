package build

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system"
)

func GitHash() string {
	return git.ShortHash(system.WorkingDirectory())
}
