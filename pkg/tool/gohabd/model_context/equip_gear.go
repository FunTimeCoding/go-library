package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) equipGear(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(constant.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	e := s.habitica.Equip(key)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	result, g := s.habitica.UserGear()

	if g != nil {
		return s.captureFail(g, constant.Unreachable)
	}

	return response.SuccessAny(convert.Gear(result))
}
