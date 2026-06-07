package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Service) ListSessions() ([]session.Session, error) {
	return s.store.ListSessions()
}
