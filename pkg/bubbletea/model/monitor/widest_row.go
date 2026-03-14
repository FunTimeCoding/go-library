package monitor

import "charm.land/bubbles/v2/table"

func widestRow(
	m *table.Model,
	index int,
) int {
	var result int

	for _, r := range m.Rows() {
		if w := len(r[index]); w > result {
			result = w
		}
	}

	return result
}
