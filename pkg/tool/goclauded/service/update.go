package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Update(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) error {
	if e := s.store.UpdateTopic(name, topic, files); e != nil {
		return e
	}

	if e := s.store.CreateCompletion(
		sessionIdentifier,
		name,
		constant.Update,
		topic,
		"",
	); e != nil {
		return e
	}

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Update,
		name,
		"",
		topic,
	); e != nil {
		return e
	}

	s.notify()

	return nil
}
