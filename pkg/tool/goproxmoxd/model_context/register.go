package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNodes,
			mcp.WithDescription("List all Proxmox nodes with status"),
		),
		mcp.NewTypedToolHandler(s.ListNodes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetNodeStatus,
			mcp.WithDescription("Get detailed resource status for a Proxmox node"),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.GetNodeStatus),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMachines,
			mcp.WithDescription(
				"List virtual machines, optionally filtered by node",
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to list across all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListMachines),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMachine,
			mcp.WithDescription(
				"Get status and config for a virtual machine by VMID",
			),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("VM ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StartMachine,
			mcp.WithDescription("Start a virtual machine"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StartMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StopMachine,
			mcp.WithDescription("Stop a virtual machine immediately"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StopMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ShutdownMachine,
			mcp.WithDescription("Gracefully shutdown a virtual machine"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ShutdownMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ResetMachine,
			mcp.WithDescription("Reset a virtual machine"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ResetMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSnapshots,
			mcp.WithDescription("List snapshots for a virtual machine"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateSnapshot,
			mcp.WithDescription("Create a snapshot of a virtual machine"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackSnapshot,
			mcp.WithDescription("Rollback a virtual machine to a snapshot"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.RollbackSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteSnapshot,
			mcp.WithDescription("Delete a snapshot from a virtual machine"),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListContainers,
			mcp.WithDescription(
				"List LXC containers, optionally filtered by node",
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to list across all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListContainers),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetContainer,
			mcp.WithDescription(
				"Get status and config for an LXC container by VMID",
			),
			mcp.WithNumber(
				"vmid",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetContainer),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNetworks,
			mcp.WithDescription(
				"List network interfaces on a Proxmox node",
			),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListNetworks),
	)
}
