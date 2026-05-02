package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SaveView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SaveView,
) (*mcp.CallToolResult, error) {
	e := s.client.SaveView(int(a.ViewIdentifier), a.FilePath)

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.Success("saved")
}
