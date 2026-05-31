package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) ContainerSnapshots(v *proxmox.Container) ([]*proxmox.ContainerSnapshot, error) {
	return v.Snapshots(c.context)
}
