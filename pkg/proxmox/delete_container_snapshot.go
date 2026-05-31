package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteContainerSnapshot(
	v *proxmox.Container,
	name string,
) (*proxmox.Task, error) {
	return v.Snapshot(name).Delete(c.context)
}
