package monitor

import "github.com/charmbracelet/bubbletea"

func addToast(m string) tea.Cmd {
	return func() tea.Msg {
		return addToastMessage(m)
	}
}
