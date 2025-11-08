package github

import "github.com/google/go-github/v78/github"

func (c *Client) PullRequests(
	owner string,
	name string,
) []*github.PullRequest {
	result, r, e := c.client.PullRequests.List(
		c.context,
		owner,
		name,
		nil,
	)
	panicOnError(r, e)

	return result
}
