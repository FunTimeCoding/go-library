package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchPages(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString("query")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("query is required: %v", f),
		), nil
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.confluence.Search(query)),
	), nil
}
