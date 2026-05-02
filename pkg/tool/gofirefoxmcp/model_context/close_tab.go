package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CloseTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CloseTab,
) (*mcp.CallToolResult, error) {
	e := s.client.CloseTab(int(a.TabIdentifier))

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.Success("closed")
}
