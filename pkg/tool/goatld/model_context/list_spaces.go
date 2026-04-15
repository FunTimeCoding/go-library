package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listSpaces(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(
		notation.MarshalIndent(s.confluence.Spaces()),
	), nil
}
