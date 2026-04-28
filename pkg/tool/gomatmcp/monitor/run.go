package monitor

func (m *Monitor) run() {
	m.refreshCache()

	for _, name := range m.configuration.Channels {
		m.processChannel(name)
	}
}
