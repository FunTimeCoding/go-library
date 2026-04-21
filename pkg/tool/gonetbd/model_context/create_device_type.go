package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDeviceType(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	model, f := r.RequireString(constant.Model)

	if f != nil {
		return response.Fail("model is required: %v", f)
	}

	manufacturer, g := r.RequireString(constant.Manufacturer)

	if g != nil {
		return response.Fail("manufacturer is required: %v", g)
	}

	m := s.client.ManufacturerByName(manufacturer)

	return response.SuccessAny(s.client.CreateDeviceType(model, m))
}
