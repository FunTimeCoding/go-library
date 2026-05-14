package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListContainers(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListContainers,
) (*mcp.CallToolResult, error) {
	var nodeNames []string

	if a.Node != "" {
		nodeNames = []string{a.Node}
	} else {
		nodes, e := s.client.Nodes()

		if e != nil {
			return s.captureDetail(e)
		}

		for _, ns := range nodes {
			nodeNames = append(nodeNames, ns.Node)
		}
	}

	var rows []proxResponse.Container

	for _, name := range nodeNames {
		node, e := s.client.Node(name)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		containers, e := s.client.Containers(node)

		if e != nil {
			return s.captureDetail(e)
		}

		for _, c := range containers {
			rows = append(
				rows,
				proxResponse.Container{
					VMID:    uint64(c.VMID),
					Name:    c.Name,
					Node:    c.Node,
					Status:  c.Status,
					CPUs:    c.CPUs,
					MaxMem:  c.MaxMem,
					MaxDisk: c.MaxDisk,
					MaxSwap: c.MaxSwap,
					Uptime:  c.Uptime,
					Tags:    c.Tags,
				},
			)
		}
	}

	return response.SuccessAny(rows)
}
