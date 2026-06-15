package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Release(
	sessionIdentifier string,
	name string,
) error {
	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Release,
		name,
		nil,
	); e != nil {
		return e
	}

	if e := s.store.ReleaseByCallsign(name); e != nil {
		return e
	}

	s.notify()

	return nil
}
