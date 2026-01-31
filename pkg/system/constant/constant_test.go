package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "src", Source)
	assert.String(t, ".osquery/shell.em", QuerySocketPath)
}
