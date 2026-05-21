package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Announce(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) {
	s.store.Announce(name, topic, files)
	s.store.LogEvent(
		sessionIdentifier,
		constant.Announce,
		name,
		"",
		topic,
	)
	s.notify()
}
