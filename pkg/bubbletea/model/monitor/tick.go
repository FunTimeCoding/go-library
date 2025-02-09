package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func tick() tea.Cmd {
	return tea.Tick(
		time.Second,
		func(t time.Time) tea.Msg {
			return TickMessage{Time: t}
		},
	)
}
