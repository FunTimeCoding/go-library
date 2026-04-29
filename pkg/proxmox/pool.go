package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Pool(name string) (*proxmox.Pool, error) {
	if name == "" {
		return nil, fmt.Errorf("pool name is required")
	}

	return c.client.Pool(c.context, name)
}
