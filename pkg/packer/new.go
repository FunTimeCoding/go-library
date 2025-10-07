package packer

import (
	"github.com/funtimecoding/go-library/pkg/packer/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func New(workDirectory string) *Client {
	web := system.Join(
		workDirectory,
		constant.PackerDirectory,
		constant.WebDirectory,
	)
	system.MakeDirectory(web)

	return &Client{
		workDirectory:      workDirectory,
		packerWebDirectory: web,
		packerOutputDirectory: system.Join(
			workDirectory,
			constant.PackerDirectory,
			constant.OutputDirectory,
		),
	}
}
