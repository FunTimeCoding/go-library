package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) scoreTask(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	direction := r.GetString(constant.Direction, "up")
	result, g := s.habitica.Score(identifier, direction)

	if g != nil {
		return s.captureFail(g, constant.Unreachable)
	}

	return response.SuccessAny(convert.ScoreResult(result))
}
