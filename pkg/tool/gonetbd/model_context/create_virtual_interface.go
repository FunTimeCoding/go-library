package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	vmName, f := r.RequireString("virtual_machine")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf(
			"virtual_machine is required: %v",
			f,
		)), nil
	}

	name, f := r.RequireString("name")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("name is required: %v", f)), nil
	}

	vm := s.client.VirtualMachineByName(vmName)

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateVirtualInterface(vm, name)),
	), nil
}
