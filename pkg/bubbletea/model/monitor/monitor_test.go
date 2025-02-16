package monitor

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMonitor(t *testing.T) {
	assert.True(t, New(false) != nil)
}
