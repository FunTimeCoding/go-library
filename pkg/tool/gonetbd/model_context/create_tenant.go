package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTenant(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	return response.SuccessAny(s.client.CreateTenant(name))
}
