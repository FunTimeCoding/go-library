package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Update(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) {
	s.store.UpdateTopic(name, topic, files)
	s.store.CreateCompletion(
		sessionIdentifier,
		name,
		constant.Update,
		topic,
		"",
	)
	s.store.LogEvent(
		sessionIdentifier,
		constant.Update,
		name,
		"",
		topic,
	)
	s.notify()
}
