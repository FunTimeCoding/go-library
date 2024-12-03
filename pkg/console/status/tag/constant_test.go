package tag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "concerns", Concerns)
	assert.String(t, "dense", Dense)
	assert.String(t, "link", Link)
	assert.String(t, "markdown", Markdown)
	assert.String(t, "score", Score)
	assert.String(t, "status", Status)
	assert.String(t, "timestamp", Timestamp)
	assert.String(t, "type", Type)
}
