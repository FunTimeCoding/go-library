package example_list

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}
