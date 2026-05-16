package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listTunnels(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.client.Tunnels()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.Tunnels(result))
}
