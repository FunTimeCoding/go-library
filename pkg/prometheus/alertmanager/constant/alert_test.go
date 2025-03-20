package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAlert(t *testing.T) {
	assert.String(t, "Container", Container)
	assert.String(t, "Deployment", Deployment)
	assert.String(t, "Node", Node)
	assert.String(t, "Pod", Pod)

	assert.String(t, "Broken", Broken)
	assert.String(t, "Corrupt", Corrupt)
	assert.String(t, "Down", Down)
	assert.String(t, "Lag", Lag)
	assert.String(t, "Load", Load)
	assert.String(t, "OutOfMemory", OutOfMemory)
	assert.String(t, "OutOfSync", OutOfSync)
	assert.String(t, "Timeout", Timeout)
}
