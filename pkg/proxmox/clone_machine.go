package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) CloneMachine(
	vm *proxmox.VirtualMachine,
	options *proxmox.VirtualMachineCloneOptions,
) (int, *proxmox.Task, error) {
	return vm.Clone(c.context, options)
}
