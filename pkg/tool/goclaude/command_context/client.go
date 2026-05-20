package command_context

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"

func (c *Context) Client() *client.ClientWithResponses {
	return c.client
}
