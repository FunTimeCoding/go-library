package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
	"log"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	var result tea.Cmd

	switch g := s.(type) {
	case tea.WindowSizeMsg:
		m.width = g.Width
		m.height = g.Height

		m.table.Columns()[0].Width = g.Width - 2

		m.table.SetHeight(g.Height - 2)
		m.table.SetWidth(g.Width - 2)
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
			log.Printf("Selected: %s\n", m.table.SelectedRow()[0])

			return m, tea.Batch(
				tea.Printf("Selected: %s", m.table.SelectedRow()[0]),
			)
		}
	}

	m.table, result = m.table.Update(s)

	return m, result
}
