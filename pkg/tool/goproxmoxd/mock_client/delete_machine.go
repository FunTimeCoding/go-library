package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteMachine(v *proxmox.VirtualMachine, _ *proxmox.VirtualMachineDeleteOptions) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	nodeMachines, okay := c.machines[v.Node]

	if !okay {
		return nil, proxmox.ErrNotFound
	}

	delete(nodeMachines, int(v.VMID))

	return &proxmox.Task{UPID: "mock:delete"}, nil
}
