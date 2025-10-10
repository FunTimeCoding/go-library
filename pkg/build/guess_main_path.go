package build

import (
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func GuessMainPath(name string) string {
	if s := join.Relative(
		constant.CommandPath,
		name,
		project.MainFile,
	); system.FileExists(s) {
		return s
	}

	return ""
}
