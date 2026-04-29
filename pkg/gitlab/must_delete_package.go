package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeletePackage(
	project int64,
	repository int64,
) {
	errors.PanicOnError(c.DeletePackage(project, repository))
}
