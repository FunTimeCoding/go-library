package monitor

func (m *Monitor) refreshCache() {
	all := m.client.MustAllChannels()
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, c := range all {
		m.channelCache[c.Name] = &c.Channel
	}
}
