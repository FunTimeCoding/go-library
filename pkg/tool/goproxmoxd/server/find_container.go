package server

import (
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/proxmox"
	upstream "github.com/luthermonson/go-proxmox"
)

func findContainer(
	c *proxmox.Client,
	identifier int64,
	node *string,
) (*upstream.Container, error) {
	if node != nil && *node != "" {
		n, e := c.Node(*node)

		if e != nil {
			return nil, e
		}

		return c.Container(n, int(identifier))
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

		containers, e := c.Containers(n)

		if e != nil {
			continue
		}

		for _, listed := range containers {
			if int64(listed.VMID) == identifier {
				return c.Container(n, int(identifier))
			}
		}
	}

	return nil, not_found.New("container", identifier)
}
