package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) removeVirtualTag(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString("virtual_machine")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf(
			"virtual_machine is required: %v",
			f,
		)), nil
	}

	tag, f := r.RequireString("tag")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("tag is required: %v", f)), nil
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.RemoveVirtualTag(name, tag)),
	), nil
}
