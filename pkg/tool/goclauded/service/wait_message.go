package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
	"time"
)

func (s *Service) WaitMessage(
	name string,
	timeout time.Duration,
) []message.Message {
	return s.store.WaitMessage(name, timeout)
}
