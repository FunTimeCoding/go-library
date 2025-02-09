package monitor

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMonitor(t *testing.T) {
	assert.True(t, New(&table.Model{}) != nil)
}
