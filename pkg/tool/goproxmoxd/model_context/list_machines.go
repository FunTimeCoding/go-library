package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMachines(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMachines,
) (*mcp.CallToolResult, error) {
	var nodeNames []string

	if a.Node != "" {
		nodeNames = []string{a.Node}
	} else {
		nodes, e := s.client.Nodes()

		if e != nil {
			return s.captureDetail(e)
		}

		for _, ns := range nodes {
			nodeNames = append(nodeNames, ns.Node)
		}
	}

	var rows []proxResponse.Machine

	for _, name := range nodeNames {
		node, e := s.client.Node(name)

		if e != nil {
			return s.captureFail(e, "node not found")
		}

		machines, e := s.client.Machines(node)

		if e != nil {
			return s.captureDetail(e)
		}

		for _, vm := range machines {
			rows = append(
				rows,
				proxResponse.Machine{
					VMID:   uint64(vm.VMID),
					Name:   vm.Name,
					Node:   vm.Node,
					Status: vm.Status,
					CPU:    vm.CPU,
					CPUs:   vm.CPUs,
					Mem:    vm.Mem,
					MaxMem: vm.MaxMem,
					Uptime: vm.Uptime,
					Tags:   vm.Tags,
				},
			)
		}
	}

	return response.SuccessAny(rows)
}
