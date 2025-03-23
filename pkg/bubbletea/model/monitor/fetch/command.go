package fetch

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Command() tea.Cmd {
	return func() tea.Msg {
		commands := []string{
			constant.GoAlert,
			constant.GoGenie,
			constant.GoSensor,
			constant.GoSentry,
			constant.GoSilence,
		}
		result := Message{}

		for _, c := range commands {
			if run.CommandExists(c) {
				m := monitor.Run(c)
				result.Items = append(result.Items, m.Items...)
			}
		}

		return result
	}
}
