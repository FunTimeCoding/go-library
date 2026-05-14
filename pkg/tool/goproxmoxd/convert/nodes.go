package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Nodes(items proxmox.NodeStatuses) []server.Node {
	result := make([]server.Node, len(items))

	for i, n := range items {
		result[i] = Node(*n)
	}

	return result
}
