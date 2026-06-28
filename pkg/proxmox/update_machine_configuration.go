package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) UpdateMachineConfiguration(
	v *proxmox.VirtualMachine,
	options ...proxmox.VirtualMachineOption,
) error {
	return v.ConfigSync(c.context, options...)
}
