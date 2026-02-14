package constant

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Equal(t, "src", Source)
	assert.Equal(t, ".osquery/shell.em", QuerySocketPath)
}
