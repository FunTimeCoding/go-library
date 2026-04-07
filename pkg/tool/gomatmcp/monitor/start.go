package monitor

import (
	"fmt"
	"log"
)

func (m *Monitor) Start() error {
	if e := m.initialize(); e != nil {
		return fmt.Errorf("initialize: %w", e)
	}

	now := nowMilli()

	for _, name := range m.configuration.Channels {
		m.lastCheckMillisecond[name] = now
	}

	if e := m.scheduler.start(); e != nil {
		return fmt.Errorf("scheduler: %w", e)
	}

	m.mutex.Lock()
	m.running = true
	m.mutex.Unlock()
	log.Printf(
		"monitor started (schedule: %s)",
		m.configuration.Schedule,
	)

	return nil
}
