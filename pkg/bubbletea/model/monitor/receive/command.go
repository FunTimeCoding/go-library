package receive

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/client"
)

func Command(c *client.Client) tea.Cmd {
	return func() tea.Msg {
		if l := c.Consume(); len(l) > 0 {
			return Message{Text: l}
		}

		return nil
	}
}
