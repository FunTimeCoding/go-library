package monitor

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
	"log"
)

func (m *Model) updateColumns() {
	remaining := m.width - 2

	for i, c := range m.table.Columns() {
		switch c.Title {
		case item.IdentifierColumn:
			title := len(item.IdentifierColumn)
			var longest int

			for _, r := range m.table.Rows() {
				if len(r[i]) > longest {
					longest = len(r[i])
				}
			}

			var width int

			if title > longest {
				width = title
			} else {
				width = longest
			}

			m.table.Columns()[i].Width = width
			remaining -= width + 2 // 2 for padding
		case item.DetailColumn:
			m.table.Columns()[i].Width = remaining - 2 // 2 for padding
			remaining = 0
		default:
			log.Panicf("unexpected: %s", c.Title)
		}
	}
}
