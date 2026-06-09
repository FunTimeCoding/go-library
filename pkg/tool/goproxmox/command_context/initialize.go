package command_context

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"

func (c *Context) Initialize(instance string) {
	c.client = client.NewEnvironment(instance)
}
