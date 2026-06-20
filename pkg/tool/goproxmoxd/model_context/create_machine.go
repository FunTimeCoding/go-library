package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument/create_machine"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a create_machine.Machine,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	identifier := a.Identifier

	if identifier <= 0 {
		identifier, e = c.NextIdentifier()

		if e != nil {
			return s.captureDetail(e)
		}
	}

	node, e := c.Node(a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	options := a.BuildOptions()
	task, e := c.CreateMachine(node, identifier, options...)

	if e != nil {
		return s.captureDetail(e)
	}

	e = c.WaitForTask(task, 300)

	if e != nil {
		return s.captureDetail(e)
	}

	if a.Start {
		vm, e := c.Machine(node, identifier)

		if e != nil {
			return s.captureDetail(e)
		}

		startTask, e := c.StartMachine(vm)

		if e != nil {
			return s.captureDetail(e)
		}

		e = c.WaitForTask(startTask, 120)

		if e != nil {
			return s.captureDetail(e)
		}
	}

	return response.SuccessAny(map[string]any{
		"identifier": identifier,
		"status":     "created",
	})
}
