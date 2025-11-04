package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Users() []*gitlab.User {
	result, r, e := c.client.Users.ListUsers(&gitlab.ListUsersOptions{})
	panicOnError(r, e)

	return result
}
