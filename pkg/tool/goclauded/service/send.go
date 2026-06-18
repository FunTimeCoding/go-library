package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
)

func (s *Service) Send(
	name string,
	to string,
	body string,
) error {
	if e := s.store.SendMessage(name, to, body); e != nil {
		return e
	}

	formatted := fmt.Sprintf("%s: %s", name, body)

	if to != "" {
		return s.PushQueue(to, constant.QueueMessage, formatted)
	}

	return s.PushQueueBroadcast(constant.QueueMessage, formatted)
}
