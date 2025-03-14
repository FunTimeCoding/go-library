package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v69/github"
)

func (c *Client) User() *github.User {
	result, _, e := c.client.Users.Get(c.context, "")
	errors.PanicOnError(e)

	return result
}
