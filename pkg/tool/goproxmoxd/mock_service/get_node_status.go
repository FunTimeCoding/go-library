package mock_service

import (
	"github.com/funtimecoding/go-library/pkg/proxmox/node_status"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
)

func (s *Service) GetNodeStatus(
	c face.ProxmoxClient,
	node string,
) (*node_status.Status, error) {
	return c.NodeStatus(node)
}
