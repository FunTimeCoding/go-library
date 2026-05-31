package model_context

import "github.com/luthermonson/go-proxmox"

func (s *Server) findMachine(
	identifier int,
	nodeName string,
) (*proxmox.VirtualMachine, error) {
	if nodeName != "" {
		node, e := s.client.Node(nodeName)

		if e != nil {
			return nil, e
		}

		return s.client.Machine(node, identifier)
	}

	nodes, e := s.client.Nodes()

	if e != nil {
		return nil, e
	}

	for _, n := range nodes {
		node, f := s.client.Node(n.Node)

		if f != nil {
			continue
		}

		machines, g := s.client.Machines(node)

		if g != nil {
			continue
		}

		for _, listed := range machines {
			if int(listed.VMID) == identifier {
				return s.client.Machine(node, identifier)
			}
		}
	}

	return nil, nil
}
