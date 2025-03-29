package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v70/github"
)

func (c *Client) Notifications() []*github.Notification {
	result, _, e := c.client.Activity.ListNotifications(
		c.context,
		&github.NotificationListOptions{},
	)
	errors.PanicOnError(e)

	return result
}
