package item

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestItem(t *testing.T) {
	assert.True(t, New([]table.Row{}) != nil)
}
