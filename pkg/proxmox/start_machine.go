package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) StartMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error) {
	return v.Start(c.context)
}
