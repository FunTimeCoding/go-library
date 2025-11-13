package gitlab

import gitlab "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Features() []*gitlab.Feature {
	result, r, e := c.client.Features.ListFeatures()
	panicOnError(r, e)

	return result
}
