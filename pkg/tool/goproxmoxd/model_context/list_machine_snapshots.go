package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMachineSnapshots(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMachineSnapshots,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	vm, e := s.findMachine(a.Identifier, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if vm == nil {
		return response.Fail("vm %d not found", a.Identifier)
	}

	snapshots, e := s.client.MachineSnapshots(vm)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(snapshots)
}
