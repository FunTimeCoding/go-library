package fetch

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Command() tea.Cmd {
	return func() tea.Msg {
		exists := run.CommandExists(constant.GoSensor)
		result := Message{}

		if exists {
			items := monitor.Run(constant.GoSensor)
			result.Items = items
		}

		return result
	}
}
