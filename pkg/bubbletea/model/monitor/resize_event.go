package monitor

import "github.com/charmbracelet/bubbletea"

func (m *Model) resizeEvent(g tea.WindowSizeMsg) {
	m.width = g.Width
	m.height = g.Height

	m.updateTableHeight(m.initialResized, len(m.toast) > 0)
	m.updateColumns()

	if !m.initialResized {
		m.initialResized = true
	}
}
