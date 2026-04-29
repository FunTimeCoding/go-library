package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Containers(n *proxmox.Node) (proxmox.Containers, error) {
	return n.Containers(c.context)
}
