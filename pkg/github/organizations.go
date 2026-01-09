package github

import "github.com/google/go-github/v81/github"

func (c *Client) Organizations(user string) []*github.Organization {
	result, r, e := c.client.Organizations.List(c.context, user, nil)
	panicOnError(r, e)

	return result
}
