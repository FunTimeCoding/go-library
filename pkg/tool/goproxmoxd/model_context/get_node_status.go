package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetNodeStatus(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetNodeStatus,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	result, e := s.client.NodeStatus(a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
