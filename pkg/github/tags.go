package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v70/github"
)

func (c *Client) Tags(
	owner string,
	repository string,
) []*github.RepositoryTag {
	result, _, e := c.client.Repositories.ListTags(
		c.context,
		owner,
		repository,
		&github.ListOptions{PerPage: 1000},
	)
	errors.PanicOnError(e)

	return result
}
