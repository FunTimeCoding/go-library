package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) CurrentUser() *gitlab.User {
	result, _, e := c.client.Users.CurrentUser()
	errors.PanicOnError(e)

	return result
}
