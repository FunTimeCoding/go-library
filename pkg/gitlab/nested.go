package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Nested() *gitlab.Client {
	return c.client
}
