package git

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func FindDirectory() string {
	return system.FindDirectoryUp(
		system.WorkDirectory(),
		constant.Directory,
	)
}
