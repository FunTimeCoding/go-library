package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"strings"
	"testing"
)

func TestWorkingDirectory(t *testing.T) {
	assert.True(
		t,
		strings.Contains(WorkingDirectory(), "system"),
	)
}
