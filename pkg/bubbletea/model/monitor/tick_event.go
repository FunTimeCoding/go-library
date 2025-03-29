package monitor

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/receive"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
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

	f := option.ExtendedColor.Copy()

	top := status.New(f).String()
	top.String(fmt.Sprintf("items: %d", len(m.table.Rows())))
	m.topBar = top.Format()

	bottom := status.New(f)
	bottom.String(m.hostname)

	if false {
		bottom.String(g.Time.Format(time.DateSecond))
	}

	bottom.Integer(60 - m.second%60)
	bottom.String(fmt.Sprintf("%dx%d", m.width, m.height))
	m.bottomBar = bottom.Format()

	m.second++
	result = append(result, tick.Command())

	if m.connect {
		result = append(result, receive.Command(m.client))
	}

	return m, tea.Batch(result...)
}
