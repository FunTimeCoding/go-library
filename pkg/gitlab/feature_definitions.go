package gitlab

import gitlab "gitlab.com/gitlab-org/api/client-go"

func (c *Client) FeatureDefinitions() []*gitlab.FeatureDefinition {
	result, r, e := c.client.Features.ListFeatureDefinitions()
	panicOnError(r, e)

	return result
}
