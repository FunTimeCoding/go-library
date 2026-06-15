package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Update(
	sessionIdentifier string,
	name string,
	topic string,
	message string,
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
		message,
	); e != nil {
		return e
	}

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Update,
		name,
		map[string]string{"topic": topic, "message": message},
	); e != nil {
		return e
	}

	s.notify()

	return nil
}
