package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
	"time"
)

func (s *Store) WaitMessage(
	name string,
	timeout time.Duration,
) ([]message.Message, error) {
	if !s.isListening(name) {
		return nil, nil
	}

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		messages, e := s.PendingMessages(name)

		if e != nil {
			return nil, e
		}

		if len(messages) > 0 {
			return messages, nil
		}

		time.Sleep(2 * time.Second)
	}

	return nil, nil
}
