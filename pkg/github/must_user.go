package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/user"
)

func (c *Client) MustUser() *user.User {
	result, e := c.User()
	errors.PanicOnError(e)

	return result
}
