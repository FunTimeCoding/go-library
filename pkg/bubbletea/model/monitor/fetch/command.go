package fetch

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"strings"
)

func Command() tea.Cmd {
	return func() tea.Msg {
		result := Message{}

		for _, c := range List() {
			if i := Run(c); len(i) > 0 {
				result.Items = append(result.Items, i...)
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
