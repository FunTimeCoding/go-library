package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) CurrentUser() *gitlab.User {
	result, r, e := c.client.Users.CurrentUser()
	panicOnError(r, e)

	return result
}
