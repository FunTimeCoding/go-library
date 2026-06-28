package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Container(
	n *proxmox.Node,
	identifier int,
) (*proxmox.Container, error) {
	nodeContainers, okay := c.containers[n.Name]

	if !okay {
		return nil, proxmox.ErrNotFound
	}

	ct, okay := nodeContainers[identifier]

	if !okay {
		return nil, proxmox.ErrNotFound
	}

	return ct, nil
}
