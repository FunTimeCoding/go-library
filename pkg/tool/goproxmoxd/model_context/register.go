package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	if len(s.service.Instances()) > 1 {
		s.server.AddTool(
			mcp.NewTool(
				constant.ListInstances,
				mcp.WithDescription(
					"List all configured Proxmox instances. Shows which instance is currently active.",
				),
			),
			mcp.NewTypedToolHandler(s.listInstances),
		)
		s.server.AddTool(
			mcp.NewTool(
				constant.UseInstance,
				mcp.WithDescription(
					"Set the active Proxmox instance for this session. Required before using any other tool.",
				),
				mcp.WithString(
					"instance",
					mcp.Required(),
					mcp.Description("Instance name from list_instances"),
				),
			),
			mcp.NewTypedToolHandler(s.useInstance),
		)
	}

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
				"identifier",
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
				"identifier",
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
				"identifier",
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
				"identifier",
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
				"identifier",
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
			constant.ListMachineSnapshots,
			mcp.WithDescription("List snapshots for a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListMachineSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMachineSnapshot,
			mcp.WithDescription("Create a snapshot of a virtual machine"),
			mcp.WithNumber(
				"identifier",
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
		mcp.NewTypedToolHandler(s.CreateMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackMachineSnapshot,
			mcp.WithDescription("Rollback a virtual machine to a snapshot"),
			mcp.WithNumber(
				"identifier",
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
		mcp.NewTypedToolHandler(s.RollbackMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteMachineSnapshot,
			mcp.WithDescription("Delete a snapshot from a virtual machine"),
			mcp.WithNumber(
				"identifier",
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
		mcp.NewTypedToolHandler(s.DeleteMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMachine,
			mcp.WithDescription(
				"Create a virtual machine. Waits for the creation task to complete before returning.",
			),
			mcp.WithString("node", mcp.Required(), mcp.Description("Target node name")),
			mcp.WithString("name", mcp.Required(), mcp.Description("VM name")),
			mcp.WithNumber(
				"identifier",
				mcp.Description("VMID (auto-allocated when omitted)"),
			),
			mcp.WithNumber("cores", mcp.Description("CPU cores (default 2)")),
			mcp.WithNumber("memory", mcp.Description("Memory in MiB (default 2048)")),
			mcp.WithNumber("disk_size", mcp.Description("Disk size in GiB (default 32)")),
			mcp.WithString(
				"disk_storage",
				mcp.Description("Storage backend (default local-lvm)"),
			),
			mcp.WithString(
				"disk_import",
				mcp.Description("Import-from path for disk image, e.g. local:import/debian-13-generic-amd64.qcow2"),
			),
			mcp.WithString("bridge", mcp.Description("Network bridge (default vnet0)")),
			mcp.WithString("ci_user", mcp.Description("Cloud-init user")),
			mcp.WithString("ci_password", mcp.Description("Cloud-init password")),
			mcp.WithString(
				"ssh_keys",
				mcp.Description("SSH public keys, newline-separated"),
			),
			mcp.WithString(
				"ip_config",
				mcp.Description("Cloud-init IP config (default ip=dhcp when cloud-init is used)"),
			),
			mcp.WithString("os_type", mcp.Description("OS type, e.g. l26 for Linux")),
			mcp.WithBoolean(
				"start",
				mcp.Description("Start VM after creation (default false)"),
			),
			mcp.WithString("tags", mcp.Description("Semicolon-separated tags")),
			mcp.WithString(
				"extras",
				mcp.Description("Additional Proxmox options as comma-separated key=value pairs"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteMachine,
			mcp.WithDescription(
				"Delete a virtual machine. The VM must be stopped first. Waits for deletion to complete.",
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
			mcp.WithBoolean(
				"purge",
				mcp.Description("Remove from cluster config and associated resources (default false)"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteMachine),
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
				"identifier",
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
			constant.ListContainerSnapshots,
			mcp.WithDescription("List snapshots for an LXC container"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListContainerSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateContainerSnapshot,
			mcp.WithDescription("Create a snapshot of an LXC container"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
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
		mcp.NewTypedToolHandler(s.CreateContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackContainerSnapshot,
			mcp.WithDescription("Rollback an LXC container to a snapshot"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
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
		mcp.NewTypedToolHandler(s.RollbackContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteContainerSnapshot,
			mcp.WithDescription("Delete a snapshot from an LXC container"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
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
		mcp.NewTypedToolHandler(s.DeleteContainerSnapshot),
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
