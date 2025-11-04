package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Option() *gitlab.Settings {
	result, r, e := c.client.Settings.GetSettings()
	panicOnError(r, e)

	return result
}
