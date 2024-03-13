package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeletePackage(
	project int,
	repository int,
) {
	_, e := c.client.Packages.DeleteProjectPackage(project, repository)
	errors.PanicOnError(e)
}
