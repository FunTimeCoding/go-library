package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualMachine(
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

	clusterName, f := r.RequireString("cluster")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"cluster is required: %v",
				f,
			),
		), nil
	}

	cl := s.client.ClusterByName(clusterName)

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateVirtualMachine(name, cl)),
	), nil
}
