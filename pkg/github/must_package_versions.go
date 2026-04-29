package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/image"
)

func (c *Client) MustPackageVersions(packageName string) []*image.Image {
	result, e := c.PackageVersions(packageName)
	errors.PanicOnError(e)

	return result
}
