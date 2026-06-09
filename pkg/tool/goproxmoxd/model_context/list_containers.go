package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListContainers(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListContainers,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	var nodeNames []string

	if a.Node != "" {
		nodeNames = []string{a.Node}
	} else {
		nodes, e := c.Nodes()

		if e != nil {
			return s.captureDetail(e)
		}

		for _, n := range nodes {
			nodeNames = append(nodeNames, n.Node)
		}
	}

	var rows []proxResponse.Container

	for _, name := range nodeNames {
		node, e := c.Node(name)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		containers, e := c.Containers(node)

		if e != nil {
			return s.captureDetail(e)
		}

		for _, ct := range containers {
			rows = append(
				rows,
				proxResponse.Container{
					Identifier: uint64(ct.VMID),
					Name:       ct.Name,
					Node:       ct.Node,
					Status:     ct.Status,
					CPUs:       ct.CPUs,
					MaxMem:     ct.MaxMem,
					MaxDisk:    ct.MaxDisk,
					MaxSwap:    ct.MaxSwap,
					Uptime:     ct.Uptime,
					Tags:       ct.Tags,
				},
			)
		}
	}

	return response.SuccessAny(rows)
}
