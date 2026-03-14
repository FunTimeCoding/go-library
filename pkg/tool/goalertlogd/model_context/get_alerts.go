package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getAlerts(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString("name")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("name is required: %v", f),
		), nil
	}

	records := s.store.ByName(name)

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Found %d alerts:\n%s",
			len(records),
			notation.MarshallIndent(records),
		),
	), nil
}
