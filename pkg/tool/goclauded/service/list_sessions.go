package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Service) ListSessions() []session.Session {
	return s.store.ListSessions()
}
