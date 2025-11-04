package github

import "github.com/funtimecoding/go-library/pkg/github/user"

func (c *Client) User() *user.User {
	result, r, e := c.client.Users.Get(c.context, "")
	panicOnError(r, e)

	return user.New(result)
}
