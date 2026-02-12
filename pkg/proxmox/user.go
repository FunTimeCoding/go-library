package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) User(name string) *proxmox.User {
	errors.PanicOnEmpty(name, constant.Name)
	result, e := c.client.User(c.context, name)
	errors.PanicOnError(e)

	return result
}
