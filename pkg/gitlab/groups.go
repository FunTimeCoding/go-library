package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Groups() []*gitlab.Group {
	result, r, e := c.client.Groups.ListGroups(nil, nil)
	panicOnError(r, e)

	return result
}
