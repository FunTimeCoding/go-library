package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"log"
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
		m.client.Write(join.Comma([]string{constant.LogoutCommand, m.user}))
		m.client.Close()

		return m, tea.Quit
	case key.Enter:
		selected := m.table.SelectedRow()
		log.Printf("Selected: %+v\n", selected)
		m.client.Write(
			join.Comma([]string{constant.FlagCommand, selected[0]}),
		)

		return m, nil
	}

	return nil, nil
}
