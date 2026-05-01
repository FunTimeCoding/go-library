package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listInterfaces(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	d, g := s.client.DeviceByNameStrict(name)

	if g != nil {
		return s.captureFail(g, "device not found")
	}

	result, h := s.client.DeviceInterfaces(d.Identifier)

	if h != nil {
		return s.captureFail(h, "interfaces not retrieved")
	}

	return response.SuccessAny(convert.Interfaces(result))
}
