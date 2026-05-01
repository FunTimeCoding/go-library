package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CloseView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CloseView,
) (*mcp.CallToolResult, error) {
	e := s.client.CloseView(int(a.ViewIdentifier))

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.Success("closed")
}
