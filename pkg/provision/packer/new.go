package packer

import (
	"github.com/funtimecoding/go-library/pkg/provision/packer/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func New(workDirectory string) *Client {
	web := join.Absolute(
		workDirectory,
		constant.PackerDirectory,
		constant.WebDirectory,
	)
	system.MakeDirectory(web)

	return &Client{
		workDirectory:      workDirectory,
		packerWebDirectory: web,
		packerOutputDirectory: join.Absolute(
			workDirectory,
			constant.PackerDirectory,
			constant.OutputDirectory,
		),
	}
}
