package monitor

func (m *Monitor) Start() {
	m.initialize()
	now := nowMilli()

	for _, name := range m.configuration.Channels {
		m.lastCheckMillisecond[name] = now
	}

	m.scheduler.Start()
	m.mutex.Lock()
	m.running = true
	m.mutex.Unlock()
	m.logger.Structured(
		"monitor_start",
		"schedule",
		m.configuration.Schedule,
		"channels",
		len(m.configuration.Channels),
	)
}
