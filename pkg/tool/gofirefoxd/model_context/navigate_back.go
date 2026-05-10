package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) NavigateBack(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.NavigateBack,
) (*mcp.CallToolResult, error) {
	e := s.client.NavigateBack(int(a.TabIdentifier))

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("navigated back")
}
