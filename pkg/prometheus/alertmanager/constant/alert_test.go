package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAlert(t *testing.T) {
	assert.String(t, "Battery", Battery)
	assert.String(t, "Cache", Cache)
	assert.String(t, "Certificate", Certificate)
	assert.String(t, "Check", Check)
	assert.String(t, "Cluster", Cluster)
	assert.String(t, "Configuration", Configuration)
	assert.String(t, "Connection", Connection)
	assert.String(t, "Container", Container)
	assert.String(t, "DaemonSet", DaemonSet)
	assert.String(t, "Database", Database)
	assert.String(t, "Date", Date)
	assert.String(t, "Deadline", Deadline)
	assert.String(t, "Deployment", Deployment)
	assert.String(t, "Disk", Disk)
	assert.String(t, "Event", Event)
	assert.String(t, "Fan", Fan)
	assert.String(t, "Hardware", Hardware)
	assert.String(t, "Latency", Latency)
	assert.String(t, "Limit", Limit)
	assert.String(t, "Measure", Measure)
	assert.String(t, "Memory", Memory)
	assert.String(t, "MountPoint", MountPoint)
	assert.String(t, "Node", Node)
	assert.String(t, "Pod", Pod)
	assert.String(t, "Queue", Queue)
	assert.String(t, "ReplicaSet", ReplicaSet)
	assert.String(t, "Replication", Replication)
	assert.String(t, "Report", Report)
	assert.String(t, "Service", Service)
	assert.String(t, "StatefulSet", StatefulSet)
	assert.String(t, "System", System)
	assert.String(t, "Token", Token)
	assert.String(t, "Volume", Volume)

	assert.String(t, "Bad", Bad)
	assert.String(t, "Behind", Behind)
	assert.String(t, "Broken", Broken)
	assert.String(t, "Close", Close)
	assert.String(t, "Corrupt", Corrupt)
	assert.String(t, "Disconnected", Disconnected)
	assert.String(t, "Down", Down)
	assert.String(t, "Empty", Empty)
	assert.String(t, "Exceeded", Exceeded)
	assert.String(t, "Expired", Expired)
	assert.String(t, "Frozen", Frozen)
	assert.String(t, "Full", Full)
	assert.String(t, "High", High)
	assert.String(t, "Inconsistent", Inconsistent)
	assert.String(t, "Lag", Lag)
	assert.String(t, "Load", Load)
	assert.String(t, "Low", Low)
	assert.String(t, "NearExceeded", NearExceeded)
	assert.String(t, "NearExpired", NearExpired)
	assert.String(t, "NearFull", NearFull)
	assert.String(t, "Old", Old)
	assert.String(t, "OutOfMemory", OutOfMemory)
	assert.String(t, "OutOfSync", OutOfSync)
	assert.String(t, "Overload", Overload)
	assert.String(t, "ReadOnly", ReadOnly)
	assert.String(t, "Slow", Slow)
	assert.String(t, "Stuck", Stuck)
	assert.String(t, "Timeout", Timeout)
	assert.String(t, "Unbound", Unbound)
	assert.String(t, "Unhealthy", Unhealthy)

	assert.String(t, "Okay", Okay)
}
