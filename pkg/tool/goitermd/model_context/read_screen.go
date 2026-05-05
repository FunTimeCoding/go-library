package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadScreen(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadScreen,
) (*mcp.CallToolResult, error) {
	v, e := s.client.ReadScreen(a.SessionIdentifier)

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.SuccessAny(v)
}
