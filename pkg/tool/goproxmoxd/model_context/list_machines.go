package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMachines(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMachines,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	var nodeNames []string

	if a.Node != "" {
		nodeNames = []string{a.Node}
	} else {
		nodes, e := c.Nodes()

		if e != nil {
			return s.captureDetail(e)
		}

		for _, n := range nodes {
			nodeNames = append(nodeNames, n.Node)
		}
	}

	var rows []proxResponse.Machine

	for _, name := range nodeNames {
		node, e := c.Node(name)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		machines, f := c.Machines(node)

		if f != nil {
			return s.captureDetail(f)
		}

		for _, m := range machines {
			rows = append(
				rows,
				proxResponse.Machine{
					Identifier: uint64(m.VMID),
					Name:       m.Name,
					Node:       m.Node,
					Status:     m.Status,
					CPU:        m.CPU,
					CPUs:       m.CPUs,
					Mem:        m.Mem,
					MaxMem:     m.MaxMem,
					Uptime:     m.Uptime,
					Tags:       m.Tags,
				},
			)
		}
	}

	return response.SuccessAny(rows)
}
