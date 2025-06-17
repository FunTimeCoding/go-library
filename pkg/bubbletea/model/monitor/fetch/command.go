package fetch

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"strings"
)

func Command() tea.Cmd {
	return func() tea.Msg {
		commands := []string{
			constant.GoAlert,
			constant.GoFile,
			constant.GoGenie,
			constant.GoJira,
			constant.GoKevt,
			constant.GoSentry,
			constant.GoSilence,
		}

		if s := environment.GetDefault(
			constant.PluginEnvironment,
			"",
		); s != "" {
			commands = append(commands, split.Comma(s)...)
		}

		result := Message{}

		for _, c := range commands {
			if run.CommandExists(c) {
				if m := monitor.Run(c); m != nil {
					result.Items = append(result.Items, m.Items...)
				}
			}
		}

		alert := fmt.Sprintf("%s-", constant.AlertPrefix)
		silence := fmt.Sprintf("%s-", constant.SilencePrefix)
		event := fmt.Sprintf("%s-", constant.KubernetesEventPrefix)
		file := fmt.Sprintf("%s-", constant.FilePrefix)

		for i, t := range result.Items {
			if strings.HasPrefix(t.Identifier, alert) {
				t.Identifier = fmt.Sprintf("%s%d", alert, i+1)

				continue
			}

			if strings.HasPrefix(t.Identifier, silence) {
				t.Identifier = fmt.Sprintf("%s%d", silence, i+1)

				continue
			}

			if strings.HasPrefix(t.Identifier, event) {
				t.Identifier = fmt.Sprintf("%s%d", event, i+1)

				continue
			}

			if strings.HasPrefix(t.Identifier, file) {
				t.Identifier = fmt.Sprintf("%s%d", file, i+1)

				continue
			}
		}

		return result
	}
}
