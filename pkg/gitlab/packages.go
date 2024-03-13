package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Packages(project int) []*gitlab.Package {
	result, _, e := c.client.Packages.ListProjectPackages(
		project,
		nil,
	)
	errors.PanicOnError(e)

	return result
}
