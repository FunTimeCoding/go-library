package github

import (
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/image"
)

func (c *Client) PackageVersions(packageName string) []*image.Image {
	result, r, e := c.client.Users.ListPackageVersions(
		c.context,
		constant.ContainerPackageType,
		packageName,
		nil,
	)
	panicOnError(r, e)

	return image.NewSlice(result)
}
