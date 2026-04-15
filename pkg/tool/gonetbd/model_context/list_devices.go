package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listDevices(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query := r.GetString("query", "")

	if query != "" {
		return mcp.NewToolResultText(
			notation.MarshalIndent(s.client.DevicesByMatch(query)),
		), nil
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.Devices()),
	), nil
}
