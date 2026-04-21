package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listDevices(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query := r.GetString(parameter.Query, "")

	if query != "" {
		return response.SuccessAny(s.client.DevicesByMatch(query))
	}

	return response.SuccessAny(s.client.Devices())
}
