package tick

import (
	"charm.land/bubbletea/v2"
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
