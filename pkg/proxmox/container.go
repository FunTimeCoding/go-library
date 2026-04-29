package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Container(
	n *proxmox.Node,
	identifier int,
) (*proxmox.Container, error) {
	return n.Container(c.context, identifier)
}
