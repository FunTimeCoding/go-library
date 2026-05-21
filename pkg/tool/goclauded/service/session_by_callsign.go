package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Service) SessionByCallsign(c string) *session.Session {
	return s.store.SessionByCallsign(c)
}
