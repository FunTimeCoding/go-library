package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func removeToast(
	identifier int,
	sleep time.Duration,
) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(sleep)

		return removeToastMessage(identifier)
	}
}
