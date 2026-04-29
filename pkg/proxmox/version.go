package proxmox

import (
	"context"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Version() (*proxmox.Version, error) {
	return c.client.Version(context.Background())
}
