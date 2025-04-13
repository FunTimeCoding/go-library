package fetch

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"strings"
)

func Command() tea.Cmd {
	return func() tea.Msg {
		commands := []string{
			constant.GoAlert,
			constant.GoGenie,
			constant.GoJira,
			constant.GoKevt,
			constant.GoSensor,
			constant.GoSentry,
			constant.GoSilence,
		}
		result := Message{}

		for _, c := range commands {
			if run.CommandExists(c) {
				if m := monitor.Run(c); m != nil {
					result.Items = append(result.Items, m.Items...)
				}
			}
		}

		for i, t := range result.Items {
			alert := fmt.Sprintf("%s-", constant.AlertPrefix)

			if strings.HasPrefix(t.Identifier, alert) {
				t.Identifier = fmt.Sprintf("%s%d", alert, i+1)
			}
		}

		return result
	}
}
