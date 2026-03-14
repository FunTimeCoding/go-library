package monitor

import "charm.land/bubbles/v2/table"

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
