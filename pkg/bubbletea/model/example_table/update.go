package example_table

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.KeyMsg:
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
			return m, tea.Batch(
				tea.Printf(
					"Let's go to %s!", m.table.SelectedRow()[1],
				),
			)
		}
	}

	table, result := m.table.Update(s)
	m.table = &table

	return m, result
}
