package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadHistory,
) (*mcp.CallToolResult, error) {
	count := int(a.Count)

	if count <= 0 {
		count = 50
	}

	v, e := s.client.ReadHistory(a.SessionIdentifier, count)

	if e != nil {
		return s.captureFail(e, constant.NotResponding)
	}

	return response.SuccessAny(v)
}
