package proxmox

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Version() *proxmox.Version {
	result, e := c.client.Version(context.Background())
	errors.PanicOnError(e)

	return result
}
