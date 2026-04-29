package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/container"
)

func (c *Client) MustPackages(user string) []*container.Container {
	result, e := c.Packages(user)
	errors.PanicOnError(e)

	return result
}
