package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) PackageFiles(
	project int,
	packageIdentifier int,
) []*gitlab.PackageFile {
	result, _, e := c.client.Packages.ListPackageFiles(
		project,
		packageIdentifier,
		nil,
	)
	errors.PanicOnError(e)

	return result
}
