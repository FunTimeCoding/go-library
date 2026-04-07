package monitor

import "log"

func (m *Monitor) run() error {
	m.refreshCache()

	for _, name := range m.configuration.Channels {
		if e := m.processChannel(name); e != nil {
			log.Printf("channel %q: %v", name, e)
		}
	}

	return nil
}
