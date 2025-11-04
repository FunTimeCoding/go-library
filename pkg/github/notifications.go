package github

import "github.com/google/go-github/v77/github"

func (c *Client) Notifications() []*github.Notification {
	result, r, e := c.client.Activity.ListNotifications(
		c.context,
		&github.NotificationListOptions{},
	)
	panicOnError(r, e)

	return result
}
