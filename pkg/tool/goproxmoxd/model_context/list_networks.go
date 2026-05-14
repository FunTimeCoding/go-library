package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListNetworks(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListNetworks,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	node, e := s.client.Node(a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	networks, e := s.client.Networks(node)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(networks)
}
