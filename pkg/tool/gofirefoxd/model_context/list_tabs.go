package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListTabs,
) (*mcp.CallToolResult, error) {
	v, e := s.client.Tabs()

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(v)
}
