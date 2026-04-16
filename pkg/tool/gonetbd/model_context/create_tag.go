package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTag(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString("name")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("name is required: %v", f)), nil
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateTag(name)),
	), nil
}
