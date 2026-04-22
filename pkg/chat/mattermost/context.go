package mattermost

import "context"

func (c *Client) Context() context.Context {
	return c.context
}
