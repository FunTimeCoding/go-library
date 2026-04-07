package monitor

import "log"

func (m *Monitor) initialize() error {
	// Resolve username for notifications
	if m.configuration.Username != "" {
		m.username = m.configuration.Username
	} else {
		me := m.client.Me()
		m.username = me.Username
	}

	log.Printf("monitor username: %s", m.username)

	// Resolve notification channel
	if m.configuration.NotificationChannel != "" {
		m.notifyChannel = m.client.TeamChannel(
			m.configuration.NotificationChannel,
		)
	} else {
		// Fall back to DM with self
		me := m.client.Me()
		m.notifyChannel = m.client.TeamChannel("town-square")
		log.Printf(
			"no notification channel configured, using town-square (me: %s)",
			me.Username,
		)
	}

	log.Printf("notify channel: %s", m.notifyChannel.Name)

	return nil
}
