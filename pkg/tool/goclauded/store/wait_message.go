package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
	"time"
)

func (s *Store) WaitMessage(
	name string,
	timeout time.Duration,
) []message.Message {
	if !s.isListening(name) {
		return nil
	}

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		messages := s.PendingMessages(name)

		if len(messages) > 0 {
			return messages
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
