package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v86/github"
)

func (c *Client) MustRepositories(owner string) []*github.Repository {
	result, e := c.Repositories(owner)
	errors.PanicOnError(e)

	return result
}
