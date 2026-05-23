package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"strings"
	"testing"
)

func TestWorkDirectory(t *testing.T) {
	assert.True(
		t,
		strings.Contains(WorkDirectory(), "system"),
	)
}
