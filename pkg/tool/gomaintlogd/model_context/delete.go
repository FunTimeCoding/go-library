package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) delete(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	id := r.GetFloat("id", 0)

	if id == 0 {
		return mcp.NewToolResultError("id is required"), nil
	}

	if e := s.store.Delete(uint(id)); e != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("failed to delete entry: %v", e),
		), nil
	}

	return mcp.NewToolResultText(
		fmt.Sprintf("Entry %d deleted successfully", int(id)),
	), nil
}
