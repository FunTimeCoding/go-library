package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Users() []*gitlab.User {
	u, _, e := c.client.Users.ListUsers(&gitlab.ListUsersOptions{})
	errors.PanicOnError(e)

	return u
}
