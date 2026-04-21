package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createAddress(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	interfaceName, g := r.RequireString(constant.Interface)

	if g != nil {
		return response.Fail("interface is required: %v", g)
	}

	address, h := r.RequireString(constant.Address)

	if h != nil {
		return response.Fail("address is required: %v", h)
	}

	d := s.client.DeviceByNameStrict(device)
	i := s.client.DeviceInterfaceByNameStrict(d, interfaceName)

	return response.SuccessAny(s.client.CreateAddress(i.Identifier, address))
}
