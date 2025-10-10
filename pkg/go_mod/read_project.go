package go_mod

import (
	"github.com/funtimecoding/go-library/pkg/go_mod/project"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func ReadProject(
	path string,
	runtimeVersion string,
) *project.Project {
	return project.New(
		path,
		ReadVersion(join.Absolute(path, ModFile)),
		runtimeVersion,
	)
}
