package monitor

import "github.com/charmbracelet/bubbles/table"

func columnWidth(
	c table.Column,
	m *table.Model,
	index int,
) int {
	column := len(c.Title)
	row := widestRow(m, index)

	if column > row {
		return column
	}

	return row
}
