package tick

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func Command() tea.Cmd {
	return tea.Tick(
		time.Second,
		func(t time.Time) tea.Msg {
			return Message{Time: t}
		},
	)
}
