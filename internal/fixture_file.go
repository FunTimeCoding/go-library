package internal

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"os"
)

func FixtureFile(path ...string) *os.File {
	return system.Open(
		system.Join(
			git.FindDirectory(),
			constant.FixturePath,
			system.Join(path...),
		),
	)
}
