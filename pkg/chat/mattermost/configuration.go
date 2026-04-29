package mattermost

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"

func (c *Client) Configuration() (map[string]string, error) {
	result, _, e := c.client.GetClientConfig(
		c.context,
		constant.EmptyEntityTag,
	)

	return result, e
}
