package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v86/github"
)

func (c *Client) MustNotifications() []*github.Notification {
	result, e := c.Notifications()
	errors.PanicOnError(e)

	return result
}
