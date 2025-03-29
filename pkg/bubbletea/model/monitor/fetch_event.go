package monitor

import "github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"

func (m *Model) fetchEvent(g fetch.Message) {
	m.items = g.Items
	rows := g.Rows()

	if !m.connect {
		for i := range rows {
			// Delete last column, user
			rows[i] = rows[i][:len(rows[i])-1]
		}
	}

	m.table.SetRows(rows)
	m.updateColumns()
}
