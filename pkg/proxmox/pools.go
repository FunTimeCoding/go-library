package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Pools() ([]*proxmox.Pool, error) {
	return c.client.Pools(c.context)
}
