package project

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func TemporaryPath(path ...string) string {
	return join.Absolute(
		git.FindDirectory(),
		constant.Temporary,
		join.Relative(path...),
	)
}
