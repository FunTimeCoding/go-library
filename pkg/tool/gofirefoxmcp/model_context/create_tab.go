package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateTab,
) (*mcp.CallToolResult, error) {
	v, e := s.client.CreateTab(a.Locator)

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.SuccessAny(v)
}
