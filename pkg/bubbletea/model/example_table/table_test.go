package example_table

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTable(t *testing.T) {
	assert.True(t, New(&table.Model{}) != nil)
}
