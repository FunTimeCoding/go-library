package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDeviceType(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	model, f := r.RequireString("model")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("model is required: %v", f)), nil
	}

	manufacturer, f := r.RequireString("manufacturer")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("manufacturer is required: %v", f)), nil
	}

	m := s.client.ManufacturerByName(manufacturer)

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateDeviceType(model, m)),
	), nil
}
