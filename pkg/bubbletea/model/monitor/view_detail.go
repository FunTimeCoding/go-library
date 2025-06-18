package monitor

import "github.com/charmbracelet/bubbletea"

func viewDetail() func() tea.Msg {
	return func() tea.Msg {
		return viewDetailMessage("")
	}
}
