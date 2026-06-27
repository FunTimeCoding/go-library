package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Machine(n *proxmox.Node, identifier int) (*proxmox.VirtualMachine, error) {
	nodeMachines, okay := c.machines[n.Name]

	if !okay {
		return nil, proxmox.ErrNotFound
	}

	vm, okay := nodeMachines[identifier]

	if !okay {
		return nil, proxmox.ErrNotFound
	}

	return vm, nil
}
