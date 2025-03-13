package monitor

import "github.com/charmbracelet/bubbles/table"

func rowIndex(
	m *table.Model,
	identifier string,
) int {
	for i, r := range m.Rows() {
		if r[0] == identifier {
			return i
		}
	}

	return -1
}
