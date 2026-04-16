package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (s *Server) createInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString("device")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"device is required: %v",
				f,
			),
		), nil
	}

	name, f := r.RequireString("name")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"name is required: %v",
				f,
			),
		), nil
	}

	interfaceType, f := r.RequireString("type")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"type is required: %v",
				f,
			),
		), nil
	}

	d := s.client.DeviceByNameStrict(device)

	return mcp.NewToolResultText(
		notation.MarshalIndent(
			s.client.CreateInterface(
				d,
				name,
				upstream.InterfaceTypeValue(interfaceType),
			),
		),
	), nil
}
