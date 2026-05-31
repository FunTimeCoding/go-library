package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goproxmoxd",
	"Proxmox hypervisor bridge for node and VM inspection",
	"goproxmoxd",
).WithInstructions(
	"Proxmox hypervisor - list and inspect nodes, virtual machines, and LXC containers.",
)

const (
	ListNodes        = "list_nodes"
	GetNodeStatus    = "get_node_status"
	ListMachines     = "list_machines"
	GetMachine       = "get_machine"
	StartMachine     = "start_machine"
	StopMachine      = "stop_machine"
	ShutdownMachine  = "shutdown_machine"
	ResetMachine     = "reset_machine"
	ListMachineSnapshots    = "list_machine_snapshots"
	CreateMachineSnapshot   = "create_machine_snapshot"
	RollbackMachineSnapshot = "rollback_machine_snapshot"
	DeleteMachineSnapshot   = "delete_machine_snapshot"
	ListContainers          = "list_containers"
	GetContainer            = "get_container"
	ListContainerSnapshots    = "list_container_snapshots"
	CreateContainerSnapshot   = "create_container_snapshot"
	RollbackContainerSnapshot = "rollback_container_snapshot"
	DeleteContainerSnapshot   = "delete_container_snapshot"
	ListNetworks     = "list_networks"
)
