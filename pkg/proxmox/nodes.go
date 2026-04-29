package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Nodes() (proxmox.NodeStatuses, error) {
	return c.client.Nodes(c.context)
}
