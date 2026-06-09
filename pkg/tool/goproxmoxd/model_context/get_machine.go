package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetMachine,
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

	if a.Node != "" {
		node, e := c.Node(a.Node)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		vm, f := c.Machine(node, a.Identifier)

		if f != nil {
			return s.captureDetail(f)
		}

		return response.SuccessAny(machineDetail(vm))
	}

	nodes, e := c.Nodes()

	if e != nil {
		return s.captureDetail(e)
	}

	for _, n := range nodes {
		node, g := c.Node(n.Node)

		if g != nil {
			return s.captureFail(g, "node not found")
		}

		machines, h := c.Machines(node)

		if h != nil {
			return s.captureDetail(h)
		}

		for _, listed := range machines {
			if uint64(listed.VMID) == uint64(a.Identifier) {
				m, i := c.Machine(node, a.Identifier)

				if i != nil {
					return s.captureDetail(i)
				}

				return response.SuccessAny(machineDetail(m))
			}
		}
	}

	return response.Fail("vm %d not found", a.Identifier)
}
