package monitor

import "github.com/charmbracelet/bubbles/table"

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
