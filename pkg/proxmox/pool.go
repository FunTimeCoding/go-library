package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Pool(name string) *proxmox.Pool {
	errors.PanicOnEmpty(name, constant.Name)
	result, e := c.client.Pool(c.context, name)
	errors.PanicOnError(e)

	return result
}
