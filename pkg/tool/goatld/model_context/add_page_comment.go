package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addPageComment(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString("identifier")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("identifier is required: %v", f),
		), nil
	}

	body, f := r.RequireString("body")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("body is required: %v", f),
		), nil
	}

	s.confluence.AddComment(identifier, body)

	return mcp.NewToolResultText("comment added"), nil
}
