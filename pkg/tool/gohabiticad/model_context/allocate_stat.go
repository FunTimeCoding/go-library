package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) allocateStat(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	stat, f := r.RequireString(constant.Stat)

	if f != nil {
		return response.Fail("stat is required: %v", f)
	}

	result, g := s.habitica.Allocate(stat)

	if g != nil {
		return s.captureFail(g, constant.Unreachable)
	}

	return response.SuccessAny(convert.Stats(result))
}
