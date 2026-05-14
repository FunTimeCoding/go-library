package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListNodes(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListNodes,
) (*mcp.CallToolResult, error) {
	nodes, e := s.client.Nodes()

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]proxResponse.Node, len(nodes))

	for i, n := range nodes {
		rows[i] = proxResponse.Node{
			Name:    n.Node,
			Status:  n.Status,
			CPU:     n.CPU,
			MaxCPU:  n.MaxCPU,
			Mem:     n.Mem,
			MaxMem:  n.MaxMem,
			Disk:    n.Disk,
			MaxDisk: n.MaxDisk,
			Uptime:  n.Uptime,
		}
	}

	return response.SuccessAny(rows)
}
