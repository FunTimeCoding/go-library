package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/tick"
)

// TODO: Begin writing the server
//  Make the client connect to the server
//  Show per item what user@host is analyzing/investigating
//  Let users flag items as important

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.WindowSizeMsg:
		m.resizeEvent(g)
	case tea.KeyMsg:
		if o, c := m.keyEvent(g); o != nil {
			return o, c
		}
	case tick.Message:
		return m.tickEvent(g)
	case fetch.Message:
		m.fetchEvent(g)
	}

	t, result := m.table.Update(s)
	m.table = &t

	return m, result
}
