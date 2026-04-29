package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Navigate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Navigate,
) (*mcp.CallToolResult, error) {
	if a.Locator == "" {
		return response.Fail("url is required")
	}

	e := s.client.Navigate(int(a.TabIdentifier), a.Locator)

	if e != nil {
		return s.captureFail(e, "navigate failed")
	}

	return response.Success("navigated to %s", a.Locator)
}
