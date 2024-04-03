package runtime

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRunningFromBinary(t *testing.T) {
	assert.False(t, RunningFromBinary())
}
