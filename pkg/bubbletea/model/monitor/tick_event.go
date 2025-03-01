package monitor

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/receive"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (m *Model) tickEvent(g tick.Message) (*Model, tea.Cmd) {
	var result tea.BatchMsg

	if m.connect && m.second == 0 {
		m.client.Connect()
		m.client.Write(
			join.Comma([]string{constant.LoginCommand, m.user, m.hostname}),
		)
	}

	if m.second%60 == 0 {
		if m.connect {
			m.client.Write(constant.PingCommand)
		}

		result = append(result, fetch.Command())
	}

	m.bottomBar = fmt.Sprintf(
		"%s %d %s",
		g.Time.Format(time.DateSecond),
		m.second,
		m.hostname,
	)
	m.second++
	result = append(result, tick.Command())

	if m.connect {
		result = append(result, receive.Command(m.client))
	}

	return m, tea.Batch(result...)
}
