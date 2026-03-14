package example_table

import "charm.land/bubbles/v2/table"

func New(m *table.Model) *Model {
	return &Model{table: m}
}
