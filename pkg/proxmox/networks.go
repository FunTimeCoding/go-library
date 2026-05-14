package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Networks(n *proxmox.Node) (proxmox.NodeNetworks, error) {
	return n.Networks(c.context)
}
