package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Groups() []*gitlab.Group {
	result, r, e := c.client.Groups.ListGroups(nil, nil)
	panicOnError(r, e)

	return result
}
