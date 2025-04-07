package tag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "age", Age)
	assert.String(t, "assignee", Assignee)
	assert.String(t, "concerns", Concerns)
	assert.String(t, "dense", Dense)
	assert.String(t, "description", Description)
	assert.String(t, "filter", Filter)
	assert.String(t, "graph", Graph)
	assert.String(t, "instance", Instance)
	assert.String(t, "investigate", Investigate)
	assert.String(t, "key", Key)
	assert.String(t, "link", Link)
	assert.String(t, "markdown", Markdown)
	assert.String(t, "score", Score)
	assert.String(t, "status", Status)
	assert.String(t, "timestamp", Timestamp)
	assert.String(t, "type", Type)
	assert.String(t, "wiki", Wiki)
}
