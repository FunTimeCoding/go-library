package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	result, g := s.confluence.Page(identifier)

	if g != nil {
		return s.captureFail(g, "page not found")
	}

	return response.SuccessAny(result)
}
