package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchIssues(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString("query")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("query is required: %v", f),
		), nil
	}

	limit := int(r.GetFloat("limit", 0))
	var results any

	if limit > 0 {
		results = s.jira.SearchLimit(limit, query)
	} else {
		results = s.jira.Search(query)
	}

	return mcp.NewToolResultText(notation.MarshalIndent(results)), nil
}
