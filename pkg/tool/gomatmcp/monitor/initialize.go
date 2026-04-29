package monitor

func (m *Monitor) initialize() {
	if m.configuration.Username != "" {
		m.username = m.configuration.Username
	} else {
		me := m.client.MustMe()
		m.username = me.Username
	}

	if m.configuration.NotificationChannel != "" {
		m.notifyChannel = m.client.MustTeamChannel(
			m.configuration.NotificationChannel,
		)
	} else {
		m.notifyChannel = m.client.MustTeamChannel("town-square")
	}

	m.logger.Structured(
		"monitor_initialize",
		"username",
		m.username,
		"notify_channel",
		m.notifyChannel.Name,
	)
}
