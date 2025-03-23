package argument

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "arm", Arm)
	assert.String(t, "cluster", Cluster)
	assert.String(t, "container", Container)
	assert.String(t, "context", Context)
	assert.String(t, "core", Core)
	assert.String(t, "disk", Disk)
	assert.String(t, "download", Download)
	assert.String(t, "file", File)
	assert.String(t, "filter", Filter)
	assert.String(t, "interactive", Interactive)
	assert.String(t, "investigate", Investigate)
	assert.String(t, "issue", Issue)
	assert.String(t, "key", Key)
	assert.String(t, "loop", Loop)
	assert.String(t, "memory", Memory)
	assert.String(t, "name", Name)
	assert.String(t, "node", Node)
	assert.String(t, "password", Password)
	assert.String(t, "pod", Pod)
	assert.String(t, "pretend", Pretend)
	assert.String(t, "retry", Retry)
	assert.String(t, "run", Run)
	assert.String(t, "summary", Summary)
	assert.String(t, "title", Title)
	assert.String(t, "topic", Topic)
}
