package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/user"
)

func (c *Client) User() *user.User {
	result, _, e := c.client.Users.Get(c.context, "")
	errors.PanicOnError(e)

	return user.New(result)
}
