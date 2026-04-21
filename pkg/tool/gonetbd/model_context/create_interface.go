package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (s *Server) createInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	name, g := r.RequireString(parameter.Name)

	if g != nil {
		return response.Fail("name is required: %v", g)
	}

	interfaceType, h := r.RequireString(constant.Type)

	if h != nil {
		return response.Fail("type is required: %v", h)
	}

	d := s.client.DeviceByNameStrict(device)

	return response.SuccessAny(
		s.client.CreateInterface(
			d,
			name,
			upstream.InterfaceTypeValue(interfaceType),
		),
	)
}
