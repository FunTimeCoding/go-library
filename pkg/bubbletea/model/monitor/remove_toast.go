package monitor

import (
	"charm.land/bubbletea/v2"
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
