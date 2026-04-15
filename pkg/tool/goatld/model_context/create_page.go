package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	space, f := r.RequireString("space_identifier")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("space_identifier is required: %v", f),
		), nil
	}

	parent, f := r.RequireString("parent_identifier")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("parent_identifier is required: %v", f),
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

	page := s.confluence.CreatePage(space, parent, title, body)

	return mcp.NewToolResultText(notation.MarshalIndent(page)), nil
}
