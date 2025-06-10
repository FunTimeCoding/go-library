package constant

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/prometheus/common/model"
)

const (
	HostEnvironment     = "PROMETHEUS_HOST"
	PortEnvironment     = "PROMETHEUS_PORT"
	UserEnvironment     = "PROMETHEUS_USER"
	PasswordEnvironment = "PROMETHEUS_PASSWORD"
)

// Metric
const (
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
	Name          = model.MetricNameLabel
	InstanceLabel = model.InstanceLabel
	JobLabel      = model.JobLabel

	ContainerLabel   = "container"
	DaemonSetLabel   = "daemonset"
	DeploymentLabel  = "deployment"
	EndpointLabel    = "endpoint"
	NamespaceLabel   = "namespace"
	NodeLabel        = "node"
	PodLabel         = "pod"
	ScopeLabel       = "scope"
	ServiceLabel     = "service"
	StateLabel       = "state"
	StatefulSetLabel = "statefulset"
)

var Format = option.Color.Copy().Tag(
	tag.Link,
	tag.Runbook,
	tag.Category,
	tag.Emoji,
)
