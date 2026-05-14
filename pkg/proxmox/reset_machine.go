package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) ResetMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error) {
	return v.Reset(c.context)
}
