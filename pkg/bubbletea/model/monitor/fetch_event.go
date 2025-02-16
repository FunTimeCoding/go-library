package monitor

import "github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"

func (m *Model) fetchEvent(g fetch.Message) {
	m.table.SetRows(g.Rows())
	m.updateColumns()
}
