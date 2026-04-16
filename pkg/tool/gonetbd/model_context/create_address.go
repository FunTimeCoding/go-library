package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createAddress(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString("device")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("device is required: %v", f)), nil
	}

	interfaceName, f := r.RequireString("interface")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("interface is required: %v", f)), nil
	}

	address, f := r.RequireString("address")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("address is required: %v", f)), nil
	}

	d := s.client.DeviceByNameStrict(device)
	i := s.client.DeviceInterfaceByNameStrict(d, interfaceName)

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateAddress(i.Identifier, address)),
	), nil
}
