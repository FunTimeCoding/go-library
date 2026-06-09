package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetContainer(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetContainer,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	if a.Node != "" {
		node, e := c.Node(a.Node)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		ct, f := c.Container(node, a.Identifier)

		if f != nil {
			return s.captureDetail(f)
		}

		return response.SuccessAny(containerDetail(ct))
	}

	nodes, e := c.Nodes()

	if e != nil {
		return s.captureDetail(e)
	}

	for _, n := range nodes {
		node, f := c.Node(n.Node)

		if f != nil {
			return s.captureFail(f, "node not found")
		}

		containers, g := c.Containers(node)

		if g != nil {
			return s.captureDetail(g)
		}

		for _, listed := range containers {
			if uint64(listed.VMID) == uint64(a.Identifier) {
				ct, h := c.Container(node, a.Identifier)

				if h != nil {
					return s.captureDetail(h)
				}

				return response.SuccessAny(containerDetail(ct))
			}
		}
	}

	return response.Fail("container %d not found", a.Identifier)
}
