package go_mod

import (
	"github.com/funtimecoding/go-library/pkg/go_mod/project"
	"github.com/funtimecoding/go-library/pkg/system"
)

func ReadProject(
	path string,
	runtimeVersion string,
) *project.Project {
	return project.New(
		path,
		ReadVersion(system.Join(path, ModFile)),
		runtimeVersion,
	)
}
