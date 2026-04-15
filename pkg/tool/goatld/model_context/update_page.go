package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updatePage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString("identifier")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("identifier is required: %v", f),
		), nil
	}

	title, f := r.RequireString("title")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("title is required: %v", f),
		), nil
	}

	body, f := r.RequireString("body")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("body is required: %v", f),
		), nil
	}

	message := r.GetString("message", "")
	page := s.confluence.UpdatePage(identifier, title, body, message)

	return mcp.NewToolResultText(notation.MarshalIndent(page)), nil
}
