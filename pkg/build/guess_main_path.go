package build

import (
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func GuessMainPath(name string) string {
	if s := system.Join(
		constant.Command,
		name,
		project.MainFile,
	); system.FileExists(s) {
		return s
	}

	return ""
}
