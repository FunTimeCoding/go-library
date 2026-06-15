package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Announce(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) error {
	if e := s.store.Announce(name, topic, files); e != nil {
		return e
	}

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Announce,
		name,
		map[string]string{"topic": topic},
	); e != nil {
		return e
	}

	s.notify()

	return nil
}
