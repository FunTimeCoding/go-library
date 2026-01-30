package github

import "github.com/google/go-github/v82/github"

func (c *Client) Repositories(owner string) []*github.Repository {
	result, r, e := c.client.Repositories.ListByUser(
		c.context,
		owner,
		nil,
	)
	panicOnError(r, e)

	return result
}
