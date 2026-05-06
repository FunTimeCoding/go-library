package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getRun(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	id := uint(r.GetInt(constant.Identifier, 0))

	if id == 0 {
		return response.Fail("id is required")
	}

	result, e := s.store.Find(id)

	if e != nil {
		return s.captureFail(e, constant.RunLookupFailed)
	}

	return response.SuccessAny(result)
}
