package monitor

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) resizeEvent(g tea.WindowSizeMsg) {
	m.width = g.Width
	m.height = g.Height
	m.table.SetHeight(m.height - 5) // top bar 1, bottom bar 1, table 3
	m.updateColumns()
}
