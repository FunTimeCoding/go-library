package monitor

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (m *Model) tickEvent(g tick.Message) (*Model, tea.Cmd) {
	var commands tea.BatchMsg

	if m.second%60 == 0 {
		commands = append(commands, fetch.Command())
	}

	m.bottomBar = fmt.Sprintf(
		"%s %d",
		g.Time.Format(time.DateSecond),
		m.second,
	)
	m.second++
	commands = append(commands, tick.Command())

	return m, tea.Batch(commands...)
}
