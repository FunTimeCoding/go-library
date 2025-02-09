package example_table

import "github.com/charmbracelet/bubbles/table"

func New(m *table.Model) *Model {
	return &Model{table: m}
}
