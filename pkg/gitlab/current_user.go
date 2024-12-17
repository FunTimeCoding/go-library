package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) CurrentUser() *gitlab.User {
	result, _, e := c.client.Users.CurrentUser()
	errors.PanicOnError(e)

	return result
}
