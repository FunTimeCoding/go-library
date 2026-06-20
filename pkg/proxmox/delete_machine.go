package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteMachine(
	v *proxmox.VirtualMachine,
	options *proxmox.VirtualMachineDeleteOptions,
) (*proxmox.Task, error) {
	return v.Delete(c.context, options)
}
