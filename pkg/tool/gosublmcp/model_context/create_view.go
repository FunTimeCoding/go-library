package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateView,
) (*mcp.CallToolResult, error) {
	v, e := s.client.CreateView(a.Title, a.Content, a.Syntax)

	if e != nil {
		return s.captureFail(e, "create view")
	}

	return response.SuccessAny(v)
}
