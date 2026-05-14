package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetContainer(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetContainer,
) (*mcp.CallToolResult, error) {
	if a.VMID == 0 {
		return response.Fail("vmid is required")
	}

	if a.Node != "" {
		node, e := s.client.Node(a.Node)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		c, e := s.client.Container(node, a.VMID)

		if e != nil {
			return s.captureDetail(e)
		}

		return response.SuccessAny(containerDetail(c))
	}

	nodes, e := s.client.Nodes()

	if e != nil {
		return s.captureDetail(e)
	}

	for _, ns := range nodes {
		node, e := s.client.Node(ns.Node)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		containers, e := s.client.Containers(node)

		if e != nil {
			return s.captureDetail(e)
		}

		for _, listed := range containers {
			if uint64(listed.VMID) == uint64(a.VMID) {
				c, e := s.client.Container(node, a.VMID)

				if e != nil {
					return s.captureDetail(e)
				}

				return response.SuccessAny(containerDetail(c))
			}
		}
	}

	return response.Fail("container %d not found", a.VMID)
}
