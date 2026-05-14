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
	ListSnapshots    = "list_snapshots"
	CreateSnapshot   = "create_snapshot"
	RollbackSnapshot = "rollback_snapshot"
	DeleteSnapshot   = "delete_snapshot"
	ListContainers   = "list_containers"
	GetContainer     = "get_container"
	ListNetworks     = "list_networks"
)
