package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/luthermonson/go-proxmox"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteMachine,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	vm, e := findMachine(c, a.Identifier, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if vm == nil {
		return response.Fail("vm %d not found", a.Identifier)
	}

	if vm.Status == "running" {
		return response.Fail(
			"vm %d is running — stop it before deleting",
			a.Identifier,
		)
	}

	task, e := c.DeleteMachine(
		vm,
		&proxmox.VirtualMachineDeleteOptions{
			Purge:                    proxmox.IntOrBool(a.Purge),
			DestroyUnreferencedDisks: proxmox.IntOrBool(true),
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	e = c.WaitForTask(task, 120)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(map[string]string{"status": "deleted"})
}
