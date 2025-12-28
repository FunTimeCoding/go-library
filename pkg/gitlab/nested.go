package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Nested() *gitlab.Client {
	return c.client
}
