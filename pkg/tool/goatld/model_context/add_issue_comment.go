package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addIssueComment(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString("key")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("key is required: %v", f),
		), nil
	}

	body, f := r.RequireString("body")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("body is required: %v", f),
		), nil
	}

	s.jira.Comment(key, body)

	return mcp.NewToolResultText("comment added"), nil
}
