package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) transitionIssue(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString("key")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("key is required: %v", f),
		), nil
	}

	transition, f := r.RequireString("transition_identifier")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("transition_identifier is required: %v", f),
		), nil
	}

	s.jira.Transition(key, transition)

	return mcp.NewToolResultText("transition applied"), nil
}
