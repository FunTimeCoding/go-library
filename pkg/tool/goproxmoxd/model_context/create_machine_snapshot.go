package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateMachineSnapshot(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateMachineSnapshot,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	vm, e := s.findMachine(a.Identifier, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if vm == nil {
		return response.Fail("vm %d not found", a.Identifier)
	}

	task, e := s.client.CreateMachineSnapshot(vm, a.Name)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(map[string]string{"task_id": string(task.UPID)})
}
