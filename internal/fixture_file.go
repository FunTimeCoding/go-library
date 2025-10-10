package internal

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"os"
)

func FixtureFile(path ...string) *os.File {
	return system.Open(
		join.Absolute(
			git.FindDirectory(),
			constant.FixturePath,
			join.Relative(path...),
		),
	)
}
