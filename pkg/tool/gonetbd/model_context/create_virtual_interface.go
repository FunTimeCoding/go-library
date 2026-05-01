package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	vmName, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	name, g := r.RequireString(parameter.Name)

	if g != nil {
		return response.Fail("name is required: %v", g)
	}

	vm, h := s.client.VirtualMachineByName(vmName)

	if h != nil {
		return s.captureFail(h, "virtual machine not found")
	}

	result, i := s.client.CreateVirtualInterface(vm, name)

	if i != nil {
		return s.captureFail(i, "virtual interface not created")
	}

	return response.SuccessAny(convert.VirtualInterface(result))
}
