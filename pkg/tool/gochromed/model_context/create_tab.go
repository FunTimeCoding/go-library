package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateTab,
) (*mcp.CallToolResult, error) {
	if a.Locator == "" {
		return response.Fail("url is required")
	}

	identifier, e := s.client.CreateTab(a.Locator)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"tab_id": identifier,
			"url":    a.Locator,
		},
	)
}
