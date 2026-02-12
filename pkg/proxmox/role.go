package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Role(name string) *proxmox.Permission {
	errors.PanicOnEmpty(name, constant.Name)
	result, e := c.client.Role(c.context, name)
	errors.PanicOnError(e)

	return &result
}
