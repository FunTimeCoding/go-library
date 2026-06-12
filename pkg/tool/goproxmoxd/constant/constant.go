package constant

import (
	"github.com/funtimecoding/go-library/pkg/identity"
	"github.com/funtimecoding/go-library/pkg/identity/paragraph"
)

var Identity = identity.New(
	"goproxmoxd",
	"Proxmox hypervisor bridge for node and VM inspection",
	"goproxmoxd",
).WithInstructions(
	"Proxmox hypervisor bridge.",
	paragraph.New(
		MultiInstance,
		"Multi-instance - call list_instances and use_instance before querying.",
	),
)

const MultiInstance = "multi_instance"
const (
	ListInstances             = "list_instances"
	UseInstance               = "use_instance"
	ListNodes                 = "list_nodes"
	GetNodeStatus             = "get_node_status"
	ListMachines              = "list_machines"
	GetMachine                = "get_machine"
	StartMachine              = "start_machine"
	StopMachine               = "stop_machine"
	ShutdownMachine           = "shutdown_machine"
	ResetMachine              = "reset_machine"
	ListMachineSnapshots      = "list_machine_snapshots"
	CreateMachineSnapshot     = "create_machine_snapshot"
	RollbackMachineSnapshot   = "rollback_machine_snapshot"
	DeleteMachineSnapshot     = "delete_machine_snapshot"
	ListContainers            = "list_containers"
	GetContainer              = "get_container"
	ListContainerSnapshots    = "list_container_snapshots"
	CreateContainerSnapshot   = "create_container_snapshot"
	RollbackContainerSnapshot = "rollback_container_snapshot"
	DeleteContainerSnapshot   = "delete_container_snapshot"
	ListNetworks              = "list_networks"
)
