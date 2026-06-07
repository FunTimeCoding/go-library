package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Complete(
	sessionIdentifier string,
	name string,
	topic string,
	message string,
) error {
	if e := s.store.UpsertCompletion(
		sessionIdentifier,
		name,
		constant.Complete,
		topic,
		message,
	); e != nil {
		return e
	}

	if e := s.store.UpsertEvent(
		sessionIdentifier,
		constant.Complete,
		name,
		topic,
		message,
	); e != nil {
		return e
	}

	s.notify()

	return nil
}
