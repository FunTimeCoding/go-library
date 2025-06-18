package monitor

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/tick"
)

func (m *Model) Init() tea.Cmd {
	return tea.Sequence(tea.EnterAltScreen, tick.Command())
}
