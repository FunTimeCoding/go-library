package constant

const KubernetesCronJobFailed = "KubernetesCronJobFailed"

// Entity
const (
	Battery       = "Battery"
	Certificate   = "Certificate"
	Check         = "Check"
	Cluster       = "Cluster"
	Configuration = "Configuration"
	Connection    = "Connection"
	Container     = "Container"
	DaemonSet     = "DaemonSet"
	Database      = "Database"
	Deployment    = "Deployment"
	Disk          = "Disk"
	Fan           = "Fan"
	Hardware      = "Hardware"
	Job           = "Job"
	Latency       = "Latency"
	Limit         = "Limit"
	Memory        = "Memory"
	Node          = "Node"
	Pod           = "Pod"
	Queue         = "Queue"
	ReplicaSet    = "ReplicaSet"
	Replication   = "Replication"
	Report        = "Report"
	Service       = "Service"
	StatefulSet   = "StatefulSet"
	Volume        = "Volume"
)

// Category
const (
	Bad          = "Bad"
	Broken       = "Broken"
	Corrupt      = "Corrupt"
	Disconnected = "Disconnected"
	Down         = "Down"
	Exceeded     = "Exceeded"
	Expired      = "Expired"
	Fail         = "Fail"
	High         = "High"
	Inconsistent = "Inconsistent"
	Lag          = "Lag"
	Load         = "Load"
	NearFull     = "NearFull"
	OutOfMemory  = "OutOfMemory"
	OutOfSync    = "OutOfSync"
	Stuck        = "Stuck"
	Timeout      = "Timeout"
	Unbound      = "Unbound"
	Unhealthy    = "Unhealthy"

	Okay = "Okay" // For Watchdog
)
