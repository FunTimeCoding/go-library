package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Complete(
	sessionIdentifier string,
	name string,
	topic string,
	message string,
) {
	s.store.UpsertCompletion(
		sessionIdentifier,
		name,
		constant.Complete,
		topic,
		message,
	)
	s.store.UpsertEvent(
		sessionIdentifier,
		constant.Complete,
		name,
		topic,
		message,
	)
	s.notify()
}
