package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
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
		if m.connect {
			m.client.Write(
				join.Comma([]string{constant.LogoutCommand, m.user}),
			)
			m.client.Close()
		}

		return m, tea.Quit
	case key.O:
		system.OpenBrowser(m.selectedItem().Link)
	case key.Enter:
		if m.connect {
			r := m.table.SelectedRow()
			m.client.Write(
				join.Comma([]string{constant.FlagCommand, r[0]}),
			)
		}

		return m, nil
	}

	return nil, nil
}
