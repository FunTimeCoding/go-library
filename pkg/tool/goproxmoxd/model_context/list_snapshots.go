package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSnapshots(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListSnapshots,
) (*mcp.CallToolResult, error) {
	if a.VMID == 0 {
		return response.Fail("vmid is required")
	}

	vm, e := s.findMachine(a.VMID, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if vm == nil {
		return response.Fail("vm %d not found", a.VMID)
	}

	snapshots, e := s.client.Snapshots(vm)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(snapshots)
}
