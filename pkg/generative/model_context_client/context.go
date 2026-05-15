package model_context_client

import "context"

func (c *Client) Context() context.Context {
	return c.context
}
