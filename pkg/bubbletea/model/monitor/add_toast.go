package monitor

import "charm.land/bubbletea/v2"

func addToast(m string) tea.Cmd {
	return func() tea.Msg {
		return addToastMessage(m)
	}
}
