package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) CloneMachine(vm *proxmox.VirtualMachine, options *proxmox.VirtualMachineCloneOptions) (int, *proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newID := options.NewID

	if newID == 0 {
		newID = c.nextID
		c.nextID++
	}

	clone := &proxmox.VirtualMachine{
		VMID: proxmox.StringOrUint64(newID),
		Name: options.Name,
		Node: vm.Node,
	}
	c.machines[vm.Node][newID] = clone

	return newID, &proxmox.Task{UPID: "mock:clone"}, nil
}
