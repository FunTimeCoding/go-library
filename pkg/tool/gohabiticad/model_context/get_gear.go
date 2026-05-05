package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getGear(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.habitica.UserGear()

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(convert.Gear(result))
}
