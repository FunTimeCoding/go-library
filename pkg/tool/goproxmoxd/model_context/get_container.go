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
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	if a.Node != "" {
		node, e := s.client.Node(a.Node)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		c, f := s.client.Container(node, a.Identifier)

		if f != nil {
			return s.captureDetail(f)
		}

		return response.SuccessAny(containerDetail(c))
	}

	nodes, e := s.client.Nodes()

	if e != nil {
		return s.captureDetail(e)
	}

	for _, n := range nodes {
		node, f := s.client.Node(n.Node)

		if f != nil {
			return s.captureFail(f, "node not found")
		}

		containers, g := s.client.Containers(node)

		if g != nil {
			return s.captureDetail(g)
		}

		for _, listed := range containers {
			if uint64(listed.VMID) == uint64(a.Identifier) {
				c, h := s.client.Container(node, a.Identifier)

				if h != nil {
					return s.captureDetail(h)
				}

				return response.SuccessAny(containerDetail(c))
			}
		}
	}

	return response.Fail("container %d not found", a.Identifier)
}
