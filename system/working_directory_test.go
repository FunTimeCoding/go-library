package system

import (
	"github.com/funtimecoding/go-library/assert"
	"strings"
	"testing"
)

func TestWorkingDirectory(t *testing.T) {
	assert.Boolean(
		t,
		true,
		strings.Contains(WorkingDirectory(), "system"),
	)
}
