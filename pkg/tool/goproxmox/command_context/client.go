package command_context

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"

func (c *Context) Client() *client.Client {
	return c.client
}
