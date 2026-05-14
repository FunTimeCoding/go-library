package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) (*proxmox.Task, error) {
	return v.DeleteSnapshot(c.context, name)
}
