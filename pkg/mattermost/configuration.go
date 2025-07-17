package mattermost

import "github.com/funtimecoding/go-library/pkg/mattermost/constant"

func (c *Client) Configuration() map[string]string {
	result, r, e := c.client.GetOldClientConfig(constant.EmptyEntityTag)
	panicOnError(e, r)

	return result
}
