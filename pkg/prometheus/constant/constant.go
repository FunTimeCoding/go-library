package constant

const (
	HostEnvironment     = "PROMETHEUS_HOST"
	PortEnvironment     = "PROMETHEUS_PORT"
	UserEnvironment     = "PROMETHEUS_USER"
	PasswordEnvironment = "PROMETHEUS_PASSWORD"
)

// Metric
const (
	Name    = "__name__"
	Up      = "up"
	Restart = "kube_pod_container_status_restarts_total"
	Load1   = "node_load1"
	Load5   = "node_load5"
	Load15  = "node_load15"
)

// Query result type
const (
	Matrix = "matrix"
	Vector = "vector"
	Scalar = "scalar"
	String = "string"
)

// Label
const (
	ContainerLabel   = "container"
	DaemonSetLabel   = "daemonset"
	DeploymentLabel  = "deployment"
	EndpointLabel    = "endpoint"
	InstanceLabel    = "instance"
	JobLabel         = "job"
	NamespaceLabel   = "namespace"
	NodeLabel        = "node"
	PodLabel         = "pod"
	ScopeLabel       = "scope"
	ServiceLabel     = "service"
	StateLabel       = "state"
	StatefulSetLabel = "statefulset"
)
