package monitor

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) resizeEvent(g tea.WindowSizeMsg) {
	m.width = g.Width
	m.height = g.Height

	m.table.Columns()[0].Width = g.Width - 2

	m.table.SetHeight(g.Height - 4)
	m.table.SetWidth(g.Width - 2)
}
