package build

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system"
)

func GitTag() string {
	return git.LatestTag(system.WorkingDirectory())
}
