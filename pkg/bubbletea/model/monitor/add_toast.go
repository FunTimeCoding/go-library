package monitor

import tea "github.com/charmbracelet/bubbletea"

func addToast(msg string) tea.Cmd {
	return func() tea.Msg {
		return addToastMessage(msg)
	}
}
