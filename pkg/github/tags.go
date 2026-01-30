package github

import "github.com/google/go-github/v82/github"

func (c *Client) Tags(
	owner string,
	repository string,
) []*github.RepositoryTag {
	result, r, e := c.client.Repositories.ListTags(
		c.context,
		owner,
		repository,
		&github.ListOptions{PerPage: 1000},
	)
	panicOnError(r, e)

	return result
}
