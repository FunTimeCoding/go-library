package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.CreateTab,
) (*mcp.CallToolResult, error) {
	v, e := s.client.CreateTab()

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.SuccessAny(v)
}
