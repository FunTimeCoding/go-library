package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Cluster() (*proxmox.Cluster, error) {
	return c.client.Cluster(c.context)
}
