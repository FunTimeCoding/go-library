package mattermost

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"

func (c *Client) Configuration() map[string]string {
	result, r, e := c.client.GetOldClientConfig(
		c.context,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
