package github

import (
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/container"
	"github.com/google/go-github/v84/github"
)

func (c *Client) Packages(user string) []*container.Container {
	result, r, e := c.client.Users.ListPackages(
		c.context,
		user,
		&github.PackageListOptions{
			PackageType: new(constant.ContainerPackageType),
		},
	)
	panicOnError(r, e)

	return container.NewSlice(result)
}
