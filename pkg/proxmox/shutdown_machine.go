package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) ShutdownMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error) {
	return v.Shutdown(c.context)
}
