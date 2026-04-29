package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetTabTitle(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetTabTitle,
) (*mcp.CallToolResult, error) {
	if a.Title == "" {
		return response.Fail("title is required")
	}

	e := s.client.SetTabTitle(a.TabIdentifier, a.Title)

	if e != nil {
		return s.captureFail(e, "set tab title")
	}

	return response.Success("set title: %s", a.Title)
}
