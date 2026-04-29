package github

import (
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/container"
	"github.com/google/go-github/v85/github"
)

func (c *Client) Packages(user string) ([]*container.Container, error) {
	result, _, e := c.client.Users.ListPackages(
		c.context,
		user,
		&github.PackageListOptions{
			PackageType: new(constant.ContainerPackageType),
		},
	)

	if e != nil {
		return nil, e
	}

	return container.NewSlice(result), nil
}
