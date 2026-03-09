package github

import "github.com/funtimecoding/go-library/pkg/github/constant"

func (c *Client) DeletePackageVersion(
	packageName string,
	versionIdentifier int64,
) {
	r, e := c.client.Users.PackageDeleteVersion(
		c.context,
		"",
		constant.ContainerPackageType,
		packageName,
		versionIdentifier,
	)
	panicOnError(r, e)
}
