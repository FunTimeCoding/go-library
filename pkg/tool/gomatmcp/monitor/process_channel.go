package monitor

import (
	"fmt"
	"log"
)

func (m *Monitor) processChannel(name string) error {
	m.mutex.RLock()
	c, okay := m.channelCache[name]
	lastMilli := m.lastCheckMillisecond[name]
	m.mutex.RUnlock()

	if !okay {
		return fmt.Errorf("not found in cache")
	}

	posts := m.client.RecentPosts(c, lastMilli)
	m.mutex.Lock()
	m.lastCheckMillisecond[name] = nowMilli()
	m.mutex.Unlock()
	relevant := findRelevantPosts(posts, m.configuration.Topics)

	if len(relevant) == 0 {
		return nil
	}

	message := buildNotification(relevant, name, m.username)

	if message == "" {
		return nil
	}

	m.client.PostSimple(m.notifyChannel, message)
	log.Printf("notified: %d posts from %s", len(relevant), name)

	return nil
}
