package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) ResizeDisk(
	v *proxmox.VirtualMachine,
	disk string,
	size string,
) (*proxmox.Task, error) {
	return v.ResizeDisk(c.context, disk, size)
}
