package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Node(name string) *proxmox.Node {
	errors.PanicOnEmpty(name, constant.Name)
	result, e := c.client.Node(c.context, name)
	errors.PanicOnError(e)

	return result
}
