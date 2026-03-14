package monitor

import "charm.land/bubbletea/v2"

func viewDetail() func() tea.Msg {
	return func() tea.Msg {
		return viewDetailMessage("")
	}
}
