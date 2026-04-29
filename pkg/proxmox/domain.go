package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Domain(name string) (*proxmox.Domain, error) {
	if name == "" {
		return nil, fmt.Errorf("domain name is required")
	}

	return c.client.Domain(c.context, name)
}
