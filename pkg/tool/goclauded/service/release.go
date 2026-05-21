package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Release(
	sessionIdentifier string,
	name string,
) {
	s.store.LogEvent(
		sessionIdentifier,
		constant.Release,
		name,
		"",
		"",
	)
	s.store.ReleaseByCallsign(name)
	s.notify()
}
