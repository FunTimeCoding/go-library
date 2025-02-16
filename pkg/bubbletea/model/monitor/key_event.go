package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
	"log"
)

func (m *Model) keyEvent(g tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch g.String() {
	case key.Escape:
		if m.table.Focused() {
			m.table.Blur()
		} else {
			m.table.Focus()
		}
	case key.Q, key.CtrlC:
		return m, tea.Quit
	case key.Enter:
		selected := m.table.SelectedRow()
		log.Printf("Selected: %+v\n", selected)

		return m, nil
	}

	return nil, nil
}
