package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"

func (s *Service) GetEvent(identifier uint) *event.Event {
	return s.store.GetEvent(identifier)
}
