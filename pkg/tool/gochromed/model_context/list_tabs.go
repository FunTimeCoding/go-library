package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListTabs,
) (*mcp.CallToolResult, error) {
	tabs := s.client.Tabs()
	type entry struct {
		Identifier string `json:"id"`
		Title      string `json:"title"`
		Locator    string `json:"url"`
	}
	var result []entry

	for _, t := range tabs {
		if t.Type != constant.PageTabType {
			continue
		}

		result = append(
			result,
			entry{
				Identifier: t.Identifier,
				Title:      t.Title,
				Locator:    t.Locator,
			},
		)
	}

	return response.SuccessAny(result)
}
