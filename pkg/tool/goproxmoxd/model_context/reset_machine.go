package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ResetMachine(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ResetMachine,
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

	task, e := s.client.ResetMachine(vm)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(map[string]string{"task_id": string(task.UPID)})
}
