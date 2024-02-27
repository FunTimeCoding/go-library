package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v59/github"
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

func (c *Client) DeleteTag(
	owner string,
	repository string,
	name string,
) {
	_, e := c.client.Git.DeleteRef(
		c.context,
		owner,
		repository,
		fmt.Sprintf("refs/tags/%s", name),
	)
	errors.PanicOnError(e)
}
