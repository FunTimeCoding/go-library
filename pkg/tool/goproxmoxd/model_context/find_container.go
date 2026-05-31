package model_context

import "github.com/luthermonson/go-proxmox"

func (s *Server) findContainer(
	identifier int,
	nodeName string,
) (*proxmox.Container, error) {
	if nodeName != "" {
		node, e := s.client.Node(nodeName)

		if e != nil {
			return nil, e
		}

		return s.client.Container(node, identifier)
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

		containers, g := s.client.Containers(node)

		if g != nil {
			continue
		}

		for _, listed := range containers {
			if int(listed.VMID) == identifier {
				return s.client.Container(node, identifier)
			}
		}
	}

	return nil, nil
}
