package project

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func FixturePath(path ...string) string {
	return join.Absolute(
		git.FindDirectory(),
		constant.FixturePath,
		join.Relative(path...),
	)
}
