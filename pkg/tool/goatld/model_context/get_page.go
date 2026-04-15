package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString("identifier")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("identifier is required: %v", f),
		), nil
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.confluence.Page(identifier)),
	), nil
}
