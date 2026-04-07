package monitor

import "log"

func (m *Monitor) Stop() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if !m.running {
		return
	}

	m.scheduler.stop()
	m.running = false
	log.Println("monitor stopped")
}
