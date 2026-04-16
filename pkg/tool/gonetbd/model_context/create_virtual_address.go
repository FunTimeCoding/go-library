package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualAddress(
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

	interfaceName, f := r.RequireString("interface")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("interface is required: %v", f)), nil
	}

	address, f := r.RequireString("address")

	if f != nil {
		return mcp.NewToolResultError(fmt.Sprintf("address is required: %v", f)), nil
	}

	vm := s.client.VirtualMachineByName(vmName)
	i := s.client.VirtualMachineInterfaceByName(vm, interfaceName)

	return mcp.NewToolResultText(
		notation.MarshalIndent(s.client.CreateVirtualAddress(i.GetId(), address)),
	), nil
}
