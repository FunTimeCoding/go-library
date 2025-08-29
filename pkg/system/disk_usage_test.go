//go:build !windows

package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	assert.NotNil(t, DiskUsage(WorkingDirectory()))
}
