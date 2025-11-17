package github

import "github.com/google/go-github/v79/github"

func (c *Client) ProjectIssues(
	owner string,
	name string,
) []*github.Issue {
	result, r, e := c.client.Issues.ListByRepo(
		c.context,
		owner,
		name,
		nil,
	)
	panicOnError(r, e)

	return result
}
