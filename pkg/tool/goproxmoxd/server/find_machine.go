package server

import (
	"github.com/funtimecoding/go-library/pkg/proxmox"
	upstream "github.com/luthermonson/go-proxmox"
)

func findMachine(
	c *proxmox.Client,
	identifier int64,
	node *string,
) (*upstream.VirtualMachine, error) {
	if node != nil && *node != "" {
		n, e := c.Node(*node)

		if e != nil {
			return nil, e
		}

		return c.Machine(n, int(identifier))
	}

	nodes, e := c.Nodes()

	if e != nil {
		return nil, e
	}

	for _, ns := range nodes {
		n, e := c.Node(ns.Node)

		if e != nil {
			continue
		}

		machines, e := c.Machines(n)

		if e != nil {
			continue
		}

		for _, listed := range machines {
			if int64(listed.VMID) == identifier {
				return c.Machine(n, int(identifier))
			}
		}
	}

	return nil, nil
}
