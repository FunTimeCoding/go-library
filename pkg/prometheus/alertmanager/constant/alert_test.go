package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAlert(t *testing.T) {
	assert.String(t, "Check", Check)
	assert.String(t, "Cluster", Cluster)
	assert.String(t, "Connection", Connection)
	assert.String(t, "Container", Container)
	assert.String(t, "Database", Database)
	assert.String(t, "Deployment", Deployment)
	assert.String(t, "Disk", Disk)
	assert.String(t, "Fan", Fan)
	assert.String(t, "Hardware", Hardware)
	assert.String(t, "Latency", Latency)
	assert.String(t, "Memory", Memory)
	assert.String(t, "Node", Node)
	assert.String(t, "Pod", Pod)
	assert.String(t, "ReplicaSet", ReplicaSet)
	assert.String(t, "Report", Report)
	assert.String(t, "Service", Service)
	assert.String(t, "Volume", Volume)

	assert.String(t, "Broken", Broken)
	assert.String(t, "Corrupt", Corrupt)
	assert.String(t, "Disconnected", Disconnected)
	assert.String(t, "Down", Down)
	assert.String(t, "High", High)
	assert.String(t, "Lag", Lag)
	assert.String(t, "Load", Load)
	assert.String(t, "NearFull", NearFull)
	assert.String(t, "OutOfMemory", OutOfMemory)
	assert.String(t, "OutOfSync", OutOfSync)
	assert.String(t, "Timeout", Timeout)
	assert.String(t, "Unhealthy", Unhealthy)
}
