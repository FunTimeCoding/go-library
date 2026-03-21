package example_list

import (
	"charm.land/bubbletea/v2"
	"fmt"
)

func (m *Model) View() tea.View {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		checked := " "

		if _, okay := m.selected[i]; okay {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	v := tea.NewView(s)
	v.WindowTitle = "Grocery List"

	return v
}
