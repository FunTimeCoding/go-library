package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) CreateContainerSnapshot(
	v *proxmox.Container,
	name string,
) (*proxmox.Task, error) {
	return v.NewSnapshot(c.context, name)
}
