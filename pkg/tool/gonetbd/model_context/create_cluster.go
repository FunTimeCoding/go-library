package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createCluster(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString("name")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"name is required: %v",
				f,
			),
		), nil
	}

	typeName, f := r.RequireString("type")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"type is required: %v",
				f,
			),
		), nil
	}

	siteName, f := r.RequireString("site")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"site is required: %v",
				f,
			),
		), nil
	}

	t := s.client.ClusterTypeByName(typeName)
	site := s.client.SiteByName(siteName)

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateCluster(name, t, site)),
	), nil
}
