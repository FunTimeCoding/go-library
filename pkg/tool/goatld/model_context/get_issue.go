package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getIssue(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString("key")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("key is required: %v", f),
		), nil
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.jira.Issue(key)),
	), nil
}
